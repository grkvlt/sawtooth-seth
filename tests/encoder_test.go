package tests

import (
    "encoding/hex"
    "testing"
    . "seth_cli/client"
    "github.com/grkvlt/sawtooth-go-sdk/signing"
)

var (
  data = []byte{0x01, 0x02, 0x03}
  PRIV_HEX = "274AAF0BFAF68C5F1CB510D2851A71CF720C7218A2C1637D635F6850FF009774"
  ENCDED="0ab40a0aca020a423033356531646533303438613632663966343738343430613232666437363535623830663061616339393762653936336231313961633534623362666465613362371280013064366530643165323133346462353833316138313964323364376132333835383836613266633063303739663837613630613130323635396237373835613532626364363037333537396537376532663964333936323136393139643134363264666430646538646136373564336233633830333066303632663032353634128001343631313764356234303934393865663265653537663061616338633932623164393831353238636561343731613230633634333963333434666230616235613334363262363830393961326665643734303532343134653034386362306632303032356266626636333736353636313365386464643164666430313164323112800137313265623534646435633965653837616561656633346164646338623765626664353231393865663165356162656466306666306132373863376262633961376439343331643531633636663236613032366336373637333162316166396335363166396131663066623065373530366464656666396162373063626663331aaf030aa4020a423033356531646533303438613632663966343738343430613232666437363535623830663061616339393762653936336231313961633534623362666465613362371a0361626322033132332a0364656632033132333a036465664a80013237383634636335323139613935316137613665353262386338646464663639383164303938646131363538643936323538633837306232633838646662636235313834316165613137326132386261666136613739373331313635353834363737303636303435633935396564306639393239363838643034646566633239524230333565316465333034386136326639663437383434306132326664373635356238306630616163393937626539363362313139616335346233626664656133623712800130643665306431653231333464623538333161383139643233643761323338353838366132666330633037396638376136306131303236353962373738356135326263643630373335373965373765326639643339363231363931396431343632646664306465386461363735643362336338303330663036326630323536341a030102031aaf030aa4020a423033356531646533303438613632663966343738343430613232666437363535623830663061616339393762653936336231313961633534623362666465613362371a0361626322033132332a0364656632033435363a036768694a80013237383634636335323139613935316137613665353262386338646464663639383164303938646131363538643936323538633837306232633838646662636235313834316165613137326132386261666136613739373331313635353834363737303636303435633935396564306639393239363838643034646566633239524230333565316465333034386136326639663437383434306132326664373635356238306630616163393937626539363362313139616335346233626664656133623712800134363131376435623430393439386566326565353766306161633863393262316439383135323863656134373161323063363433396333343466623061623561333436326236383039396132666564373430353234313465303438636230663230303235626662663633373635363631336538646464316466643031316432311a03010203"
)

func TestEncoder(t *testing.T) {
    private_key, _ := hex.DecodeString(PRIV_HEX)
    encoder := NewEncoder(private_key, TransactionParams{
        FamilyName: "abc",
        FamilyVersion: "123",
        Inputs: []string{"def"},
    })

    txn1 := encoder.NewTransaction(data, TransactionParams{
        Nonce: "123",
        Outputs: []string{"def"},
    })

    priv := signing.NewSecp256k1PrivateKey(private_key)
    pubstr := signing.NewSecp256k1Context().GetPublicKey(priv).AsHex()
    txn2 := encoder.NewTransaction(data, TransactionParams{
        Nonce: "456",
        Outputs: []string{"ghi"},
        BatcherPublicKey: pubstr,
    })

    // Test serialization
    txns, err := ParseTransactions(SerializeTransactions([]*Transaction{txn1, txn2}))
    if err != nil {
        t.Error(err)
    }

    batch := encoder.NewBatch(txns)

    // Test serialization
    batches, err := ParseBatches(SerializeBatches([]*Batch{batch}))
    if err != nil {
        t.Error(err)
    }
    data := SerializeBatches(batches)
    datastr := hex.EncodeToString(data)

    expected := ENCDED

    if datastr != expected {
        t.Error("Did not correctly encode batch. Got", datastr)
    }
}
