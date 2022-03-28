package guardianChecker

import "errors"

// ErrAccountHasNoGuardianSet signals that the account has no guardians set
var ErrAccountHasNoGuardianSet = errors.New("account has no guardian set")

// ErrActiveHasNoActiveGuardian signals that the account has no active guardian
var ErrActiveHasNoActiveGuardian = errors.New("account has no active guardian")

// ErrNilEpochNotifier signals that a nil epoch notifier was provided
var ErrNilEpochNotifier = errors.New("nil epoch notifier provided")
