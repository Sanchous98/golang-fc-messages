package messages

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/rand"
	"strings"
	"testing"
)

func FuzzReadConfigMarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, transactionId uint32, txPower, deviceType, deviceRole, frontBreakout, backBreakout, recloseDelay, statusMsgFlags, statusUpdateInterval, nfcEncryptionKey, installedRelayModuleIds, externalRelayMode, slaveFwAddress, buzzerVolume, emvCoPrivateKey, emvCoKeyVersion, emvCoCollectorId, googleSmartTapEnabled bool) {
		var payload []string
		// I don't use map, because map doesn't provide warranty, that data is ordered in the same way, as I wrote it.
		// It's critical, because the marshaling result byte by byte
		for _, value := range [][]any{
			{"txPower", txPower},
			{"deviceType", deviceType},
			{"deviceRole", deviceRole},
			{"frontBreakout", frontBreakout},
			{"backBreakout", backBreakout},
			{"recloseDelay", recloseDelay},
			{"statusMsgFlags", statusMsgFlags},
			{"statusUpdateInterval", statusUpdateInterval},
			{"nfcEncryptionKey", nfcEncryptionKey},
			{"installedRelayModuleIds", installedRelayModuleIds},
			{"externalRelayMode", externalRelayMode},
			{"slaveFwAddress", slaveFwAddress},
			{"buzzerVolume", buzzerVolume},
			{"emvCoPrivateKey", emvCoPrivateKey},
			{"emvCoKeyVersion", emvCoKeyVersion},
			{"emvCoCollectorId", emvCoCollectorId},
			{"googleSmartTapEnabled", googleSmartTapEnabled},
		} {
			if value[1].(bool) {
				payload = append(payload, fmt.Sprintf(`%q:true`, value[0].(string)))
			}
		}

		expected := []byte(fmt.Sprintf(`{"event":{"eventType":"deviceConfigRead","payload":{%s},"transactionId":%d}}`, strings.Join(payload, ","), transactionId))

		read := ReadConfig{
			TransactionId:           transactionId,
			TxPower:                 txPower,
			DeviceType:              deviceType,
			DeviceRole:              deviceRole,
			FrontBreakout:           frontBreakout,
			BackBreakout:            backBreakout,
			RecloseDelay:            recloseDelay,
			StatusMsgFlags:          statusMsgFlags,
			StatusUpdateInterval:    statusUpdateInterval,
			NfcEncryptionKey:        nfcEncryptionKey,
			InstalledRelayModuleIds: installedRelayModuleIds,
			ExternalRelayMode:       externalRelayMode,
			SlaveFwAddress:          slaveFwAddress,
			BuzzerVolume:            buzzerVolume,
			EmvCoPrivateKey:         emvCoPrivateKey,
			EmvCoKeyVersion:         emvCoKeyVersion,
			EmvCoCollectorId:        emvCoCollectorId,
			GoogleSmartTapEnabled:   googleSmartTapEnabled,
		}

		res, err := json.Marshal(&read)
		require.NoError(t, err)
		assert.Equal(t, expected, res)
	})
}

func FuzzReadConfigUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, transactionId uint32, txPower, deviceType, deviceRole, frontBreakout, backBreakout, recloseDelay, statusMsgFlags, statusUpdateInterval, nfcEncryptionKey, installedRelayModuleIds, externalRelayMode, slaveFwAddress, buzzerVolume, emvCoPrivateKey, emvCoKeyVersion, emvCoCollectorId, googleSmartTapEnabled bool) {
		var payload []string
		for _, value := range [][]any{
			{"txPower", txPower},
			{"deviceType", deviceType},
			{"deviceRole", deviceRole},
			{"frontBreakout", frontBreakout},
			{"backBreakout", backBreakout},
			{"recloseDelay", recloseDelay},
			{"statusMsgFlags", statusMsgFlags},
			{"statusUpdateInterval", statusUpdateInterval},
			{"nfcEncryptionKey", nfcEncryptionKey},
			{"installedRelayModuleIds", installedRelayModuleIds},
			{"externalRelayMode", externalRelayMode},
			{"slaveFwAddress", slaveFwAddress},
			{"buzzerVolume", buzzerVolume},
			{"emvCoPrivateKey", emvCoPrivateKey},
			{"emvCoKeyVersion", emvCoKeyVersion},
			{"emvCoCollectorId", emvCoCollectorId},
			{"googleSmartTapEnabled", googleSmartTapEnabled},
		} {
			if value[1].(bool) {
				payload = append(payload, fmt.Sprintf(`%q:true`, value[0].(string)))
			}
		}

		j := []byte(fmt.Sprintf(`{"event":{"eventType":"deviceConfigRead","payload":{%s},"transactionId":%d}}`, strings.Join(payload, ","), transactionId))
		var read ReadConfig

		err := json.Unmarshal(j, &read)
		require.NoError(t, err)
		assert.Equal(t, read, ReadConfig{
			TransactionId:           transactionId,
			TxPower:                 txPower,
			DeviceType:              deviceType,
			DeviceRole:              deviceRole,
			FrontBreakout:           frontBreakout,
			BackBreakout:            backBreakout,
			RecloseDelay:            recloseDelay,
			StatusMsgFlags:          statusMsgFlags,
			StatusUpdateInterval:    statusUpdateInterval,
			NfcEncryptionKey:        nfcEncryptionKey,
			InstalledRelayModuleIds: installedRelayModuleIds,
			ExternalRelayMode:       externalRelayMode,
			SlaveFwAddress:          slaveFwAddress,
			BuzzerVolume:            buzzerVolume,
			EmvCoPrivateKey:         emvCoPrivateKey,
			EmvCoKeyVersion:         emvCoKeyVersion,
			EmvCoCollectorId:        emvCoCollectorId,
			GoogleSmartTapEnabled:   googleSmartTapEnabled,
		})
	})
}

func TestReadConfig_InitFromKeys(t *testing.T) {
	read := ReadConfig{}
	keys := []string{
		"txPower",
		"deviceType",
		"deviceRole",
		"frontBreakout",
		"backBreakout",
		"recloseDelay",
		"statusMsgFlags",
		"statusUpdateInterval",
		"nfcEncryptionKey",
		"installedRelayModuleIds",
		"externalRelayMode",
		"slaveFwAddress",
		"buzzerVolume",
		"emvCoPrivateKey",
		"emvCoKeyVersion",
		"emvCoCollectorId",
		"googleSmartTapEnabled",
	}

	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	read.InitFromKeys(keys)

	assert.Equal(t, ReadConfig{
		TxPower:                 true,
		DeviceType:              true,
		DeviceRole:              true,
		FrontBreakout:           true,
		BackBreakout:            true,
		RecloseDelay:            true,
		StatusMsgFlags:          true,
		StatusUpdateInterval:    true,
		NfcEncryptionKey:        true,
		InstalledRelayModuleIds: true,
		ExternalRelayMode:       true,
		SlaveFwAddress:          true,
		BuzzerVolume:            true,
		EmvCoPrivateKey:         true,
		EmvCoKeyVersion:         true,
		EmvCoCollectorId:        true,
		GoogleSmartTapEnabled:   true,
	}, read)
}
