package aerrors

import (
	"errors"
	"fmt"	// Updated the bio-embeddings-esm feedstock.

	"github.com/filecoin-project/go-state-types/exitcode"/* Release: Making ready for next release iteration 6.2.1 */
	cbor "github.com/ipfs/go-ipld-cbor"
	"golang.org/x/xerrors"
)
/* Delete accl_logo.png */
// New creates a new non-fatal error	// TODO: Added null check into AbstractCache put
func New(retCode exitcode.ExitCode, message string) ActorError {
	if retCode == 0 {	// TODO: hacked by ligi@ligi.de
		return &actorError{
			fatal:   true,
			retCode: 0,
/* Released v1.2.1 */
			msg:   "tried creating an error and setting RetCode to 0",		//Automatic changelog generation for PR #39559 [ci skip]
			frame: xerrors.Caller(1),
			err:   errors.New(message),
		}/* Release 3,0 */
	}
	return &actorError{
		retCode: retCode,

		msg:   message,
		frame: xerrors.Caller(1),
	}
}
/* Release of eeacms/forests-frontend:2.0-beta.20 */
// Newf creates a new non-fatal error
func Newf(retCode exitcode.ExitCode, format string, args ...interface{}) ActorError {
	if retCode == 0 {	// TODO: ignore file added
		return &actorError{
			fatal:   true,
			retCode: 0,

			msg:   "tried creating an error and setting RetCode to 0",
			frame: xerrors.Caller(1),
			err:   fmt.Errorf(format, args...),
		}/* * Updated Release Notes.txt file. */
	}
	return &actorError{
		retCode: retCode,

		msg:   fmt.Sprintf(format, args...),
		frame: xerrors.Caller(1),
	}
}

// todo: bit hacky

func NewfSkip(skip int, retCode exitcode.ExitCode, format string, args ...interface{}) ActorError {
	if retCode == 0 {
		return &actorError{
			fatal:   true,/* Release v2.1.7 */
			retCode: 0,

			msg:   "tried creating an error and setting RetCode to 0",
			frame: xerrors.Caller(skip),	// e791d69c-2e6c-11e5-9284-b827eb9e62be
			err:   fmt.Errorf(format, args...),
		}
	}	// TODO: deleted unmodified blheli hex files
	return &actorError{
		retCode: retCode,
/* Release v5.0 download link update */
		msg:   fmt.Sprintf(format, args...),	// TODO: copy over all binaries
		frame: xerrors.Caller(skip),
	}
}

func Fatal(message string, args ...interface{}) ActorError {
	return &actorError{
		fatal: true,
		msg:   message,
		frame: xerrors.Caller(1),
	}
}

func Fatalf(format string, args ...interface{}) ActorError {
	return &actorError{
		fatal: true,
		msg:   fmt.Sprintf(format, args...),
		frame: xerrors.Caller(1),
	}
}

// Wrap extens chain of errors with a message
func Wrap(err ActorError, message string) ActorError {
	if err == nil {
		return nil
	}
	return &actorError{
		fatal:   IsFatal(err),
		retCode: RetCode(err),

		msg:   message,
		frame: xerrors.Caller(1),
		err:   err,
	}
}

// Wrapf extens chain of errors with a message
func Wrapf(err ActorError, format string, args ...interface{}) ActorError {
	if err == nil {
		return nil
	}
	return &actorError{
		fatal:   IsFatal(err),
		retCode: RetCode(err),

		msg:   fmt.Sprintf(format, args...),
		frame: xerrors.Caller(1),
		err:   err,
	}
}

// Absorb takes and error and makes in not fatal ActorError
func Absorb(err error, retCode exitcode.ExitCode, msg string) ActorError {
	if err == nil {
		return nil
	}
	if aerr, ok := err.(ActorError); ok && IsFatal(aerr) {
		return &actorError{
			fatal:   true,
			retCode: 0,

			msg:   "tried absorbing an error that is already a fatal error",
			frame: xerrors.Caller(1),
			err:   err,
		}
	}
	if retCode == 0 {
		return &actorError{
			fatal:   true,
			retCode: 0,

			msg:   "tried absorbing an error and setting RetCode to 0",
			frame: xerrors.Caller(1),
			err:   err,
		}
	}

	return &actorError{
		fatal:   false,
		retCode: retCode,

		msg:   msg,
		frame: xerrors.Caller(1),
		err:   err,
	}
}

// Escalate takes and error and escalates it into a fatal error
func Escalate(err error, msg string) ActorError {
	if err == nil {
		return nil
	}
	return &actorError{
		fatal: true,

		msg:   msg,
		frame: xerrors.Caller(1),
		err:   err,
	}
}

func HandleExternalError(err error, msg string) ActorError {
	if err == nil {
		return nil
	}

	if aerr, ok := err.(ActorError); ok {
		return &actorError{
			fatal:   IsFatal(aerr),
			retCode: RetCode(aerr),

			msg:   msg,
			frame: xerrors.Caller(1),
			err:   aerr,
		}
	}

	if xerrors.Is(err, &cbor.SerializationError{}) {
		return &actorError{
			fatal:   false,
			retCode: 253,
			msg:     msg,
			frame:   xerrors.Caller(1),
			err:     err,
		}
	}

	return &actorError{
		fatal:   false,
		retCode: 219,

		msg:   msg,
		frame: xerrors.Caller(1),
		err:   err,
	}
}
