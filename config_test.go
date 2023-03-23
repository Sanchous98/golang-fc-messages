package messages

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/rand"
	"reflect"
	"strings"
	"testing"
)

func FuzzReadConfigMarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, transactionId int, txPower, deviceType, deviceRole, frontBreakout, backBreakout, recloseDelay, statusMsgFlags, statusUpdateInterval, nfcEncryptionKey, installedRelayModuleIds, externalRelayMode, slaveFwAddress, buzzerVolume, emvCoPrivateKey, emvCoKeyVersion, emvCoCollectorId, googleSmartTapEnabled bool) {
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
				payload = append(payload, fmt.Sprintf(`"%s":true`, value[0].(string)))
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
	f.Fuzz(func(t *testing.T, transactionId int, txPower, deviceType, deviceRole, frontBreakout, backBreakout, recloseDelay, statusMsgFlags, statusUpdateInterval, nfcEncryptionKey, installedRelayModuleIds, externalRelayMode, slaveFwAddress, buzzerVolume, emvCoPrivateKey, emvCoKeyVersion, emvCoCollectorId, googleSmartTapEnabled bool) {
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
				payload = append(payload, fmt.Sprintf(`"%s":true`, value[0].(string)))
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

func FuzzUpdateConfigMarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, transactionId int, txPower uint, deviceType, deviceRole, frontBreakout, backBreakout string, recloseDelay, statusMsgFlags uint, statusUpdateInterval uint16, nfcEncryptionKey []byte, externalRelayMode string, slaveFwAddress uint, buzzerVolume string, emvCoPrivateKey, emvCoKeyVersion, emvCoCollectorId []byte, googleSmartTapEnabled bool) {
		var installedRelayModuleIds [16]uint

		for index := range installedRelayModuleIds {
			installedRelayModuleIds[index] = uint(rand.Int())
		}

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
			{"nfcEncryptionKey", [16]byte(append(make([]byte, 16), nfcEncryptionKey...))},
			{"installedRelayModuleIds", installedRelayModuleIds},
			{"externalRelayMode", externalRelayMode},
			{"slaveFwAddress", slaveFwAddress},
			{"buzzerVolume", buzzerVolume},
			{"emvCoPrivateKey", [32]byte(append(make([]byte, 32), emvCoPrivateKey...))},
			{"emvCoKeyVersion", [4]byte(append(make([]byte, 4), emvCoKeyVersion...))},
			{"emvCoCollectorId", [4]byte(append(make([]byte, 4), emvCoCollectorId...))},
			{"googleSmartTapEnabled", googleSmartTapEnabled},
		} {
			if !reflect.ValueOf(value[1]).IsZero() || reflect.ValueOf(value[1]).Kind() == reflect.Array {
				value[1], _ = json.Marshal(value[1])

				payload = append(payload, fmt.Sprintf(`"%s":%s`, value[0], value[1]))
			}
		}

		expected := []byte(fmt.Sprintf(`{"event":{"eventType":"deviceConfigUpdate","payload":{%s},"transactionId":%d}}`, strings.Join(payload, ","), transactionId))

		update := UpdateConfig{
			TransactionId:           transactionId,
			TxPower:                 txPower,
			DeviceType:              deviceType,
			DeviceRole:              deviceRole,
			FrontBreakout:           frontBreakout,
			BackBreakout:            backBreakout,
			RecloseDelay:            recloseDelay,
			StatusMsgFlags:          statusMsgFlags,
			StatusUpdateInterval:    statusUpdateInterval,
			NfcEncryptionKey:        [16]byte(append(make([]byte, 16), nfcEncryptionKey...)),
			InstalledRelayModuleIds: installedRelayModuleIds,
			ExternalRelayMode:       externalRelayMode,
			SlaveFwAddress:          slaveFwAddress,
			BuzzerVolume:            buzzerVolume,
			EmvCoPrivateKey:         [32]byte(append(make([]byte, 32), emvCoPrivateKey...)),
			EmvCoKeyVersion:         [4]byte(append(make([]byte, 4), emvCoKeyVersion...)),
			EmvCoCollectorId:        [4]byte(append(make([]byte, 4), emvCoCollectorId...)),
			GoogleSmartTapEnabled:   googleSmartTapEnabled,
		}

		res, err := json.Marshal(&update)

		require.NoError(t, err)
		assert.Equal(t, expected, res)
	})
}

func FuzzUpdateConfigUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, transactionId int, txPower uint, deviceType, deviceRole, frontBreakout, backBreakout string, recloseDelay, statusMsgFlags uint, statusUpdateInterval uint16, nfcEncryptionKey []byte, externalRelayMode string, slaveFwAddress uint, buzzerVolume string, emvCoPrivateKey, emvCoKeyVersion, emvCoCollectorId []byte, googleSmartTapEnabled bool) {
		var installedRelayModuleIds [16]uint

		for index := range installedRelayModuleIds {
			installedRelayModuleIds[index] = uint(rand.Int())
		}

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
			{"nfcEncryptionKey", [16]byte(append(make([]byte, 16), nfcEncryptionKey...))},
			{"installedRelayModuleIds", installedRelayModuleIds},
			{"externalRelayMode", externalRelayMode},
			{"slaveFwAddress", slaveFwAddress},
			{"buzzerVolume", buzzerVolume},
			{"emvCoPrivateKey", [32]byte(append(make([]byte, 32), emvCoPrivateKey...))},
			{"emvCoKeyVersion", [4]byte(append(make([]byte, 4), emvCoKeyVersion...))},
			{"emvCoCollectorId", [4]byte(append(make([]byte, 4), emvCoCollectorId...))},
			{"googleSmartTapEnabled", googleSmartTapEnabled},
		} {
			if !reflect.ValueOf(value[1]).IsZero() || reflect.ValueOf(value[1]).Kind() == reflect.Array {
				switch value[1].(type) {
				case string:
					value[1], _ = json.Marshal(value[1].(string))
				default:
					value[1], _ = json.Marshal(value[1])
				}

				payload = append(payload, fmt.Sprintf(`"%s":%s`, value[0], value[1]))
			}
		}

		value := []byte(fmt.Sprintf(`{"event":{"eventType":"deviceConfigUpdate","payload":{%s},"transactionId":%d}}`, strings.Join(payload, ","), transactionId))

		var res UpdateConfig

		err := json.Unmarshal(value, &res)

		expected := UpdateConfig{
			TransactionId:           transactionId,
			TxPower:                 txPower,
			DeviceType:              deviceType,
			DeviceRole:              deviceRole,
			FrontBreakout:           frontBreakout,
			BackBreakout:            backBreakout,
			RecloseDelay:            recloseDelay,
			StatusMsgFlags:          statusMsgFlags,
			StatusUpdateInterval:    statusUpdateInterval,
			NfcEncryptionKey:        [16]byte(append(make([]byte, 16), nfcEncryptionKey...)),
			InstalledRelayModuleIds: installedRelayModuleIds,
			ExternalRelayMode:       externalRelayMode,
			SlaveFwAddress:          slaveFwAddress,
			BuzzerVolume:            buzzerVolume,
			EmvCoPrivateKey:         [32]byte(append(make([]byte, 32), emvCoPrivateKey...)),
			EmvCoKeyVersion:         [4]byte(append(make([]byte, 4), emvCoKeyVersion...)),
			EmvCoCollectorId:        [4]byte(append(make([]byte, 4), emvCoCollectorId...)),
			GoogleSmartTapEnabled:   googleSmartTapEnabled,
		}
		require.NoError(t, err)
		assert.Equal(t, expected, res)
	})
}
