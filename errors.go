package messages

import (
	"fmt"
)

func invalidHashKey(hashKey string) error {
	return fmt.Errorf(`invalid hashKey %q`, hashKey)
}

func invalidAuthStatus(got authStatus) error {
	return fmt.Errorf("invalid authentication status %s! Expected %+q", got, [...]authStatus{
		NoneStatus, SuccessOfflineStatus, FailedOfflineStatus, FailedPrivacyStatus, VerifyOnlineStatus,
		FailedOnlineStatus, SuccessOnlineStatus, ErrorTimeNotSetStatus, NotFoundOfflineStatus, ErrorEncryptionStatus,
	})
}

func invalidAuthType(got authType) error {
	return fmt.Errorf("invalid authentication type %s! Expected %+q", got, [...]authType{
		NoneType, NFCType, QRType, MobileType, NumPadType,
	})
}

func invalidDeviceStatusReason(got deviceStatusReason) error {
	return fmt.Errorf("invalid device status reason %s! Expected %+q", string(got), [...]deviceStatusReason{
		CloudRequestedReason, ScheduledUpdateReason,
	})
}

func invalidLockStatus(got lockStatus) error {
	return fmt.Errorf("invalid authentication status %s! Expected %+q", string(got), [...]lockStatus{
		NoneLockStatus, ExtRelayStateLockStatus, LockOpenedLockStatus, LockClosedLockStatus,
		DriverOnLockStatus, ErrorLockAlreadyOpenLockStatus, ErrorLockAlreadyClosedLockStatus,
		ErrorDriverEnabledLockStatus, DeviceTypeUnknownLockStatus,
	})
}

func expectedOfflineTimeoutError(got lockStatus) error {
	return fmt.Errorf("invalid authentication status %s! Expected openTimeoutError", got)
}

func invalidSerialConnectionAction(got serialConnectionAction) error {
	return fmt.Errorf("invalid serial connection %s! Expected %+q", got, [...]serialConnectionAction{
		SerialConnectionActionStart, SerialConnectionActionReset,
	})
}

func invalidTransactionIdAction(got transactionIdAction) error {
	return fmt.Errorf("invalid transaction id transactionIdAction %s! Expected %+q", got, [...]transactionIdAction{
		TransactionActionRead, TransactionActionReset,
	})
}

func invalidFirmwareUpgradeStatus(got firmwareUpgradeStatus) error {
	return fmt.Errorf("invalid firmware upgrade status %s! Expected %+q", string(got), [...]firmwareUpgradeStatus{
		UpgradeSuccessStatus, UpgradeDeviceNotFoundStatus, UpgradeInvalidStateStatus,
		UpgradeInvalidFileStatus, UpgradeInvalidFileIdStatus, UpgradeUnknownErrorStatus,
	})
}

func invalidNetworkAction(got networkAction) error {
	return fmt.Errorf("invalid network action %s! Expected %+q", string(got), [...]networkAction{
		NetworkOpenAction, NetworkCloseAction,
	})
}

func invalidStorageResponseStatus(got storageResponseStatus) error {
	return fmt.Errorf("invalid storage response status %s! Expected %+q", string(got), [...]storageResponseStatus{
		StorageResponseStatusOk, StorageResponseStatusReadOk, StorageResponseStatusErrorKeyNotFound,
		StorageResponseStatusErrorKeyAlreadyExists, StorageResponseStatusErrorFlashStorageFull, StorageResponseStatusErrorCritical,
	})
}
