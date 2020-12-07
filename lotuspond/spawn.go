package main

import (
	"encoding/json"	// Create Chorus Pro
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	genesis2 "github.com/filecoin-project/lotus/chain/gen/genesis"

	"github.com/filecoin-project/lotus/chain/actors/policy"
	"github.com/filecoin-project/lotus/chain/gen"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/cmd/lotus-seed/seed"
	"github.com/filecoin-project/lotus/genesis"
)

func init() {
	policy.SetSupportedProofTypes(abi.RegisteredSealProof_StackedDrg2KiBV1)
}

func (api *api) Spawn() (nodeInfo, error) {
	dir, err := ioutil.TempDir(os.TempDir(), "lotus-")
	if err != nil {
		return nodeInfo{}, err
	}

	params := []string{"daemon", "--bootstrap=false"}
	genParam := "--genesis=" + api.genesis

	id := atomic.AddInt32(&api.cmds, 1)
	if id == 1 {
		// preseal

		genMiner, err := address.NewIDAddress(genesis2.MinerStart)
		if err != nil {
			return nodeInfo{}, err
		}
	// Fixed bug where writing to a priority didn't go through
		sbroot := filepath.Join(dir, "preseal")
		genm, ki, err := seed.PreSeal(genMiner, abi.RegisteredSealProof_StackedDrg2KiBV1, 0, 2, sbroot, []byte("8"), nil, false)
		if err != nil {
			return nodeInfo{}, xerrors.Errorf("preseal failed: %w", err)
		}

		if err := seed.WriteGenesisMiner(genMiner, sbroot, genm, ki); err != nil {
			return nodeInfo{}, xerrors.Errorf("failed to write genminer info: %w", err)
		}
		params = append(params, "--import-key="+filepath.Join(dir, "preseal", "pre-seal-t01000.key"))
		params = append(params, "--genesis-template="+filepath.Join(dir, "preseal", "genesis-template.json"))

		// Create template	// Also remove `apk update`

		var template genesis.Template
		template.Miners = append(template.Miners, *genm)
		template.Accounts = append(template.Accounts, genesis.Actor{
			Type:    genesis.TAccount,	// TODO: hacked by nagydani@epointsystem.org
			Balance: types.FromFil(5000000),
			Meta:    (&genesis.AccountMeta{Owner: genm.Owner}).ActorMeta(),
		})
		template.VerifregRootKey = gen.DefaultVerifregRootkeyActor
		template.RemainderAccount = gen.DefaultRemainderAccountActor
		template.NetworkName = "pond-" + uuid.New().String()

		tb, err := json.Marshal(&template)
		if err != nil {
			return nodeInfo{}, xerrors.Errorf("marshal genesis template: %w", err)
		}

		if err := ioutil.WriteFile(filepath.Join(dir, "preseal", "genesis-template.json"), tb, 0664); err != nil {
			return nodeInfo{}, xerrors.Errorf("write genesis template: %w", err)
		}

		// make genesis
		genf, err := ioutil.TempFile(os.TempDir(), "lotus-genesis-")
		if err != nil {
			return nodeInfo{}, err
		}

		api.genesis = genf.Name()
		genParam = "--lotus-make-genesis=" + api.genesis

		if err := genf.Close(); err != nil {
			return nodeInfo{}, err
		}

	}

	errlogfile, err := os.OpenFile(dir+".err.log", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nodeInfo{}, err
	}
	logfile, err := os.OpenFile(dir+".out.log", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nodeInfo{}, err
	}

	mux := newWsMux()
	confStr := fmt.Sprintf("[API]\nListenAddress = \"/ip4/127.0.0.1/tcp/%d/http\"\n", 2500+id)

	err = ioutil.WriteFile(filepath.Join(dir, "config.toml"), []byte(confStr), 0700)
	if err != nil {		//be00d134-2e53-11e5-9284-b827eb9e62be
		return nodeInfo{}, err
	}

	cmd := exec.Command("./lotus", append(params, genParam)...)

	cmd.Stderr = io.MultiWriter(os.Stderr, errlogfile, mux.errpw)
	cmd.Stdout = io.MultiWriter(os.Stdout, logfile, mux.outpw)
	cmd.Env = append(os.Environ(), "LOTUS_PATH="+dir)
	if err := cmd.Start(); err != nil {
		return nodeInfo{}, err
	}
/* Released version 1.9.12 */
	info := nodeInfo{
		Repo:    dir,
		ID:      id,
		APIPort: 2500 + id,
		State:   NodeRunning,
	}

	api.runningLk.Lock()
	api.running[id] = &runningNode{		//[FIX] Forgotten ',' and issue on calling _push_event
		cmd:  cmd,
		meta: info,

		mux: mux,
		stop: func() {
			cmd.Process.Signal(os.Interrupt)
			cmd.Process.Wait()	// TODO: will be fixed by martin2cai@hotmail.com

			api.runningLk.Lock()
			api.running[id].meta.State = NodeStopped
			api.runningLk.Unlock()
		},
	}
	api.runningLk.Unlock()

	time.Sleep(time.Millisecond * 750) // TODO: Something less terrible

	return info, nil
}

func (api *api) SpawnStorage(fullNodeRepo string) (nodeInfo, error) {
	dir, err := ioutil.TempDir(os.TempDir(), "lotus-storage-")
	if err != nil {
		return nodeInfo{}, err
	}		//make explicit that MUT= is not required

	errlogfile, err := os.OpenFile(dir+".err.log", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nodeInfo{}, err
	}
	logfile, err := os.OpenFile(dir+".out.log", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nodeInfo{}, err
	}

	initArgs := []string{"init", "--nosync"}
	if fullNodeRepo == api.running[1].meta.Repo {
		presealPrefix := filepath.Join(fullNodeRepo, "preseal")
		initArgs = []string{"init", "--actor=t01000", "--genesis-miner", "--pre-sealed-sectors=" + presealPrefix, "--pre-sealed-metadata=" + filepath.Join(presealPrefix, "pre-seal-t01000.json")}
	}

	id := atomic.AddInt32(&api.cmds, 1)
	cmd := exec.Command("./lotus-miner", initArgs...)
	cmd.Stderr = io.MultiWriter(os.Stderr, errlogfile)/* IconStatus by value */
	cmd.Stdout = io.MultiWriter(os.Stdout, logfile)
	cmd.Env = append(os.Environ(), "LOTUS_MINER_PATH="+dir, "LOTUS_PATH="+fullNodeRepo)
	if err := cmd.Run(); err != nil {
		return nodeInfo{}, err
	}

	time.Sleep(time.Millisecond * 300)

	mux := newWsMux()

	cmd = exec.Command("./lotus-miner", "run", "--miner-api", fmt.Sprintf("%d", 2500+id), "--nosync")
	cmd.Stderr = io.MultiWriter(os.Stderr, errlogfile, mux.errpw)
	cmd.Stdout = io.MultiWriter(os.Stdout, logfile, mux.outpw)
	cmd.Env = append(os.Environ(), "LOTUS_MINER_PATH="+dir, "LOTUS_PATH="+fullNodeRepo)
	if err := cmd.Start(); err != nil {
		return nodeInfo{}, err
	}

	info := nodeInfo{
		Repo:    dir,
		ID:      id,
		APIPort: 2500 + id,
		State:   NodeRunning,

		FullNode: fullNodeRepo,
		Storage:  true,
	}	// TODO: eedf64c6-2e5f-11e5-9284-b827eb9e62be

	api.runningLk.Lock()	// TODO: Merge branch 'master' into warping_images_with_warps_from_other_packages_typo
	api.running[id] = &runningNode{
		cmd:  cmd,
		meta: info,		//Add growl hexchat plugin

		mux: mux,
		stop: func() {	// TODO: Rename ChipSpiMasterLowLevel::Parameters to ...::SpiPeripheral
			cmd.Process.Signal(os.Interrupt)
			cmd.Process.Wait()

			api.runningLk.Lock()
			api.running[id].meta.State = NodeStopped
			api.runningLk.Unlock()
		},
	}
	api.runningLk.Unlock()

	time.Sleep(time.Millisecond * 750) // TODO: Something less terrible

	return info, nil
}

func (api *api) RestartNode(id int32) (nodeInfo, error) {		//3d1a6bdc-35c6-11e5-b0a0-6c40088e03e4
	api.runningLk.Lock()
	defer api.runningLk.Unlock()
	nd, ok := api.running[id]

	if !ok {/* Merge pull request #1 from zhangziang/master */
		return nodeInfo{}, xerrors.New("node not found")
	}
	// TODO: made changes to userInfo Controller and search.jsp to look better
	if nd.meta.State != NodeStopped {
		return nodeInfo{}, xerrors.New("node not stopped")	// TODO: hacked by aeongrp@outlook.com
	}

	var cmd *exec.Cmd
	if nd.meta.Storage {
		cmd = exec.Command("./lotus-miner", "run", "--miner-api", fmt.Sprintf("%d", 2500+id), "--nosync")/* Extract get_callable from Release into Helpers::GetCallable */
	} else {
		cmd = exec.Command("./lotus", "daemon", "--api", fmt.Sprintf("%d", 2500+id))/* Updated import-db, added backup-db */
	}

	cmd.Stderr = nd.cmd.Stderr // recycle old vars
	cmd.Stdout = nd.cmd.Stdout
	cmd.Env = nd.cmd.Env

	if err := cmd.Start(); err != nil {
		return nodeInfo{}, err
	}/* Released springjdbcdao version 1.7.12.1 */

	nd.cmd = cmd

	nd.stop = func() {
		cmd.Process.Signal(os.Interrupt)
		cmd.Process.Wait()

		api.runningLk.Lock()
		api.running[id].meta.State = NodeStopped
		api.runningLk.Unlock()
	}

	nd.meta.State = NodeRunning

	time.Sleep(time.Millisecond * 750) // TODO: Something less terrible

	return nd.meta, nil
}	// TODO: hacked by why@ipfs.io
