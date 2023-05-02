package messages

import "fmt"

type InvalidEventType struct {
	Got eventType
}

func (e InvalidEventType) Error() string { return "invalid event type " + string(e.Got) }

type InvalidHashKey struct {
	HashKey string
}

func (e InvalidHashKey) Error() string { return "invalid hashKey " + e.HashKey }

type InvalidAuthStatus struct {
	Got authStatus
}

func (e InvalidAuthStatus) Error() string {
	return fmt.Sprintf("invalid authentication status %s! Expected %+q", e.Got, [...]authStatus{
		NoneStatus, SuccessOfflineStatus, FailedOfflineStatus, FailedPrivacyStatus, VerifyOnlineStatus,
		FailedOnlineStatus, SuccessOnlineStatus, ErrorTimeNotSetStatus, NotFoundOfflineStatus, ErrorEncryptionStatus,
	})
}

type InvalidAuthType struct {
	Got authType
}

func (e InvalidAuthType) Error() string {
	return fmt.Sprintf("invalid authentication type %s! Expected %+q", e.Got, [...]authType{
		NoneType, NFCType, QRType, MobileType, NumPadType,
	})
}

type InvalidDeviceStatusReason struct {
	Got deviceStatusReason
}

func (e InvalidDeviceStatusReason) Error() string {
	return fmt.Sprintf("invalid device status reason %s! Expected %+q", e.Got, [...]deviceStatusReason{
		CloudRequestedReason, ScheduledUpdateReason,
	})
}

type InvalidLockStatus struct {
	Got lockStatus
}

func (e InvalidLockStatus) Error() string {
	return fmt.Sprintf("invalid authentication status %s! Expected %+q", e.Got, [...]lockStatus{
		NoneLockStatus, ExtRelayStateLockStatus, LockOpenedLockStatus, LockClosedLockStatus,
		DriverOnLockStatus, ErrorLockAlreadyOpenLockStatus, ErrorLockAlreadyClosedLockStatus,
		ErrorDriverEnabledLockStatus, DeviceTypeUnknownLockStatus,
	})
}

type ExpectedOfflineTimeoutError struct {
	Got lockStatus
}

func (e ExpectedOfflineTimeoutError) Error() string {
	return fmt.Sprintf("invalid authentication status %s! Expected openTimeoutError", e.Got)
}

type InvalidSerialConnectionAction struct {
	Got serialConnectionAction
}

func (e InvalidSerialConnectionAction) Error() string {
	return fmt.Sprintf("invalid serial connection %s! Expected %+q", e.Got, [...]serialConnectionAction{
		SerialConnectionActionStart, SerialConnectionActionReset,
	})
}

type InvalidTransactionIdAction struct {
	Got transactionIdAction
}

func (e InvalidTransactionIdAction) Error() string {
	return fmt.Sprintf(`invalid transaction id transactionIdAction %s! Expected %+q`, e.Got, [...]transactionIdAction{
		TransactionActionRead, TransactionActionReset,
	})
}

type InvalidFirmwareUpgradeStatus struct {
	Got firmwareUpgradeStatus
}

func (e InvalidFirmwareUpgradeStatus) Error() string {
	return fmt.Sprintf("invalid firmware upgrade status %s! Expected %+q", e.Got, [...]firmwareUpgradeStatus{
		UpgradeSuccessStatus, UpgradeDeviceNotFoundStatus, UpgradeInvalidStateStatus,
		UpgradeInvalidFileStatus, UpgradeInvalidFileIdStatus, UpgradeUnknownErrorStatus,
	})
}

type InvalidNetworkAction struct {
	Got networkAction
}

func (e InvalidNetworkAction) Error() string {
	return fmt.Sprintf("invalid network action %s! Expected %+q", e.Got, [...]networkAction{
		NetworkOpenAction, NetworkCloseAction,
	})
}

type InvalidStorageResponseStatus struct {
	Got storageResponseStatus
}

func (e InvalidStorageResponseStatus) Error() string {
	return fmt.Sprintf("invalid storage response status %d! Expected %+q", e.Got, [...]storageResponseStatus{
		StorageResponseStatusOk, StorageResponseStatusReadOk, StorageResponseStatusErrorKeyNotFound,
		StorageResponseStatusErrorKeyAlreadyExists, StorageResponseStatusErrorFlashStorageFull, StorageResponseStatusErrorCritical,
	})
}

type InvalidConfigResponseStatus struct {
	Got configResponseStatus
}

func (e InvalidConfigResponseStatus) Error() string {
	return fmt.Sprintf("invalid config response status \"%s\"! Expected %+q", e.Got, [...]configResponseStatus{
		ResponseStatusNone, ResponseStatusCreateOK, ResponseStatusReadOK, ResponseStatusUpdateOK,
		ResponseStatusDeleteOK, ResponseStatusConfigSizeError, ResponseStatusError, ResponseStatusErrorOutOfRange,
		ResponseStatusErrorNotFound, ResponseStatusErrorFlash, ResponseStatusErrorNoCallBack, ResponseStatusErrorNoSpace,
		ResponseStatusErrorNoReadAccess, ResponseStatusErrorNoWriteAccess,
	})
}
