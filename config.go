package messages

import (
	"github.com/goccy/go-json"
)

const (
	DeviceConfigReadEvent     eventType = "deviceConfigRead"
	DeviceConfigUpdateEvent   eventType = "deviceConfigUpdate"
	DeviceConfigResponseEvent eventType = "deviceConfigResponse"
)

const (
	ResponseStatusNone               configResponseStatus = "none"
	ResponseStatusCreateOK           configResponseStatus = "createOK"
	ResponseStatusReadOK             configResponseStatus = "readOK"
	ResponseStatusUpdateOK           configResponseStatus = "updateOK"
	ResponseStatusDeleteOK           configResponseStatus = "deleteOK"
	ResponseStatusConfigSizeError    configResponseStatus = "configSizeError"
	ResponseStatusError              configResponseStatus = "error"
	ResponseStatusErrorOutOfRange    configResponseStatus = "errorOutOfRange"
	ResponseStatusErrorNotFound      configResponseStatus = "errorNotFound"
	ResponseStatusErrorFlash         configResponseStatus = "errorFlash"
	ResponseStatusErrorNoCallBack    configResponseStatus = "errorNoCallBack"
	ResponseStatusErrorNoSpace       configResponseStatus = "errorNoSpace"
	ResponseStatusErrorNoReadAccess  configResponseStatus = "errorNoReadAccess"
	ResponseStatusErrorNoWriteAccess configResponseStatus = "errorNoWriteAccess"
)

type configResponseStatus string

type UpdateConfig struct {
	TransactionId uint32  `json:"-"`
	TxPower       *uint   `json:"txPower,omitempty"`
	DeviceType    *string `json:"deviceType,omitempty"`
	DeviceRole    *string `json:"deviceRole,omitempty"`
	//FrontBreakout           string   `json:"frontBreakout,omitempty"`
	//BackBreakout            string   `json:"backBreakout,omitempty"`
	RecloseDelay            *uint   `json:"recloseDelay,omitempty"`
	StatusMsgFlags          *uint   `json:"statusMsgFlags,omitempty"`
	StatusUpdateInterval    *uint16 `json:"statusUpdateInterval,omitempty"`
	NfcPiccEncryptionKey    *string `json:"nfcPiccEncryptionKey,omitempty"`
	NfcEncryptionKey        *string `json:"nfcEncryptionKey,omitempty"`
	InstalledRelayModuleIds []uint  `json:"installedRelayModuleIds,omitempty"`
	ExternalRelayMode       *string `json:"externalRelayMode,omitempty"`
	SlaveFwAddress          *uint   `json:"slaveFwAddress,omitempty"`
	BuzzerVolume            *string `json:"buzzerVolume,omitempty"`
	EmvCoPrivateKey         *string `json:"emvCoPrivateKey,omitempty"`
	EmvCoKeyVersion         *string `json:"emvCoKeyVersion,omitempty"`
	EmvCoCollectorId        *string `json:"emvCoCollectorId,omitempty"`
	GoogleSmartTapEnabled   *bool   `json:"googleSmartTapEnabled,omitempty"`
}

func (r *UpdateConfig) UnmarshalJSON(bytes []byte) error {
	type updateConfig UpdateConfig

	var e event
	var err error

	if err = json.Unmarshal(bytes, &e); err != nil {
		return err
	}

	if e.EventType != DeviceConfigUpdateEvent {
		return e.EventType.Error()
	}

	if err = json.Unmarshal(e.Payload, (*updateConfig)(r)); err != nil {
		return err
	}

	r.TransactionId = e.TransactionId

	return nil
}

func (r *UpdateConfig) MarshalJSON() ([]byte, error) {
	type updateConfig UpdateConfig

	var e event
	var err error

	e.TransactionId = r.TransactionId
	e.EventType = DeviceConfigUpdateEvent

	if e.Payload, err = json.Marshal((*updateConfig)(r)); err != nil {
		return nil, err
	}

	return json.Marshal(&e)
}

type ConfigResponse struct {
	ShortAddr            string               `json:"-"`
	ExtAddr              string               `json:"-"`
	Rssi                 int                  `json:"-"`
	TransactionId        uint32               `json:"-"`
	Status               configResponseStatus `json:"status"`
	TxPower              *uint                `json:"txPower,omitempty"`
	DeviceType           *string              `json:"deviceType,omitempty"`
	DeviceRole           *string              `json:"deviceRole,omitempty"`
	FrontBreakout        *string              `json:"frontBreakout,omitempty"`
	BackBreakout         *string              `json:"backBreakout,omitempty"`
	RecloseDelay         *uint                `json:"recloseDelay,omitempty"`
	StatusMsgFlags       *uint                `json:"statusMsgFlags,omitempty"`
	StatusUpdateInterval *uint16              `json:"statusUpdateInterval,omitempty"`
	//NfcPiccEncryptionKey    [16]byte `json:"nfcPiccEncryptionKey,omitempty"`
	//NfcEncryptionKey        [16]byte `json:"nfcEncryptionKey,omitempty"`
	InstalledRelayModuleIds []uint  `json:"installedRelayModuleIds,omitempty"`
	ExternalRelayMode       *string `json:"externalRelayMode,omitempty"`
	SlaveFwAddress          *uint   `json:"slaveFwAddress,omitempty"`
	BuzzerVolume            *string `json:"buzzerVolume,omitempty"`
	//EmvCoPrivateKey         string `json:"emvCoPrivateKey,omitempty"`
	EmvCoKeyVersion       *string `json:"emvCoKeyVersion,omitempty"`
	EmvCoCollectorId      *string `json:"emvCoCollectorId,omitempty"`
	GoogleSmartTapEnabled *bool   `json:"googleSmartTapEnabled,omitempty"`
}

func (r *ConfigResponse) UnmarshalJSON(bytes []byte) error {
	type configResponse ConfigResponse

	var e response
	var err error

	if err = json.Unmarshal(bytes, &e); err != nil {
		return err
	}

	if e.EventType != DeviceConfigResponseEvent {
		return e.EventType.Error()
	}

	if err = json.Unmarshal(e.Payload, (*configResponse)(r)); err != nil {
		return err
	}

	switch r.Status {
	case ResponseStatusNone, ResponseStatusCreateOK, ResponseStatusReadOK, ResponseStatusUpdateOK, ResponseStatusDeleteOK,
		ResponseStatusConfigSizeError, ResponseStatusError, ResponseStatusErrorOutOfRange, ResponseStatusErrorNotFound,
		ResponseStatusErrorFlash, ResponseStatusErrorNoCallBack, ResponseStatusErrorNoSpace, ResponseStatusErrorNoReadAccess,
		ResponseStatusErrorNoWriteAccess:
	default:
		return InvalidConfigResponseStatus{r.Status}
	}

	r.TransactionId = e.TransactionId
	r.ShortAddr = e.ShortAddr
	r.ExtAddr = e.ExtAddr
	r.Rssi = e.Rssi

	return nil
}

func (r *ConfigResponse) MarshalJSON() ([]byte, error) {
	switch r.Status {
	case ResponseStatusNone, ResponseStatusCreateOK, ResponseStatusReadOK, ResponseStatusUpdateOK, ResponseStatusDeleteOK,
		ResponseStatusConfigSizeError, ResponseStatusError, ResponseStatusErrorOutOfRange, ResponseStatusErrorNotFound,
		ResponseStatusErrorFlash, ResponseStatusErrorNoCallBack, ResponseStatusErrorNoSpace, ResponseStatusErrorNoReadAccess,
		ResponseStatusErrorNoWriteAccess:
	default:
		return nil, InvalidConfigResponseStatus{r.Status}
	}

	type configResponse ConfigResponse

	var e response
	var err error

	e.TransactionId = r.TransactionId
	e.ShortAddr = r.ShortAddr
	e.ExtAddr = r.ExtAddr
	e.Rssi = r.Rssi
	e.EventType = DeviceConfigResponseEvent

	if e.Payload, err = json.Marshal((*configResponse)(r)); err != nil {
		return nil, err
	}

	return json.Marshal(&e)
}

type ReadConfig struct {
	TxPower                 bool   `json:"txPower,omitempty"`
	DeviceType              bool   `json:"deviceType,omitempty"`
	DeviceRole              bool   `json:"deviceRole,omitempty"`
	FrontBreakout           bool   `json:"frontBreakout,omitempty"`
	BackBreakout            bool   `json:"backBreakout,omitempty"`
	RecloseDelay            bool   `json:"recloseDelay,omitempty"`
	StatusMsgFlags          bool   `json:"statusMsgFlags,omitempty"`
	StatusUpdateInterval    bool   `json:"statusUpdateInterval,omitempty"`
	NfcPiccEncryptionKey    bool   `json:"nfcPiccEncryptionKey,omitempty"`
	NfcEncryptionKey        bool   `json:"nfcEncryptionKey,omitempty"`
	InstalledRelayModuleIds bool   `json:"installedRelayModuleIds,omitempty"`
	ExternalRelayMode       bool   `json:"externalRelayMode,omitempty"`
	SlaveFwAddress          bool   `json:"slaveFwAddress,omitempty"`
	BuzzerVolume            bool   `json:"buzzerVolume,omitempty"`
	EmvCoPrivateKey         bool   `json:"emvCoPrivateKey,omitempty"`
	EmvCoKeyVersion         bool   `json:"emvCoKeyVersion,omitempty"`
	EmvCoCollectorId        bool   `json:"emvCoCollectorId,omitempty"`
	GoogleSmartTapEnabled   bool   `json:"googleSmartTapEnabled,omitempty"`
	TransactionId           uint32 `json:"-"`
}

func (r *ReadConfig) InitFromKeys(keys []string) *ReadConfig {
	for _, key := range keys {
		switch key {
		case "txPower":
			r.TxPower = true
		case "deviceType":
			r.DeviceType = true
		case "deviceRole":
			r.DeviceRole = true
		case "frontBreakout":
			r.FrontBreakout = true
		case "backBreakout":
			r.BackBreakout = true
		case "recloseDelay":
			r.RecloseDelay = true
		case "statusMsgFlags":
			r.StatusMsgFlags = true
		case "statusUpdateInterval":
			r.StatusUpdateInterval = true
		case "nfcEncryptionKey":
			r.NfcEncryptionKey = true
		case "nfcPiccEncryptionKey":
			r.NfcPiccEncryptionKey = true
		case "installedRelayModuleIds":
			r.InstalledRelayModuleIds = true
		case "externalRelayMode":
			r.ExternalRelayMode = true
		case "slaveFwAddress":
			r.SlaveFwAddress = true
		case "buzzerVolume":
			r.BuzzerVolume = true
		case "emvCoPrivateKey":
			r.EmvCoPrivateKey = true
		case "emvCoKeyVersion":
			r.EmvCoKeyVersion = true
		case "emvCoCollectorId":
			r.EmvCoCollectorId = true
		case "googleSmartTapEnabled":
			r.GoogleSmartTapEnabled = true
		}
	}

	return r
}

func (r *ReadConfig) UnmarshalJSON(bytes []byte) error {
	type readConfig ReadConfig

	var e event
	var err error

	if err = json.Unmarshal(bytes, &e); err != nil {
		return err
	}

	if e.EventType != DeviceConfigReadEvent {
		return e.EventType.Error()
	}

	if err = json.Unmarshal(e.Payload, (*readConfig)(r)); err != nil {
		return err
	}

	r.TransactionId = e.TransactionId

	return nil
}

func (r *ReadConfig) MarshalJSON() ([]byte, error) {
	type readConfig ReadConfig

	var e event
	var err error

	e.TransactionId = r.TransactionId
	e.EventType = DeviceConfigReadEvent

	if e.Payload, err = json.Marshal((*readConfig)(r)); err != nil {
		return nil, err
	}

	return json.Marshal(&e)
}
