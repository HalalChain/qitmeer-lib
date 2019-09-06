package message

import (
	"github.com/Qitmeer/qitmeer-lib/core/dag"
	"io"
)

type MsgGraphState struct {
	GS *dag.GraphState
}

func (msg *MsgGraphState) Decode(r io.Reader, pver uint32) error {
	msg.GS=dag.NewGraphState()
	err:=msg.GS.Decode(r,pver)
	if err != nil {
		return err
	}
	return nil
}

func (msg *MsgGraphState) Encode(w io.Writer, pver uint32) error {
	err := msg.GS.Encode(w,pver)
	if err != nil {
		return err
	}
	return nil
}

func (msg *MsgGraphState) Command() string {
	return CmdGraphState
}

func (msg *MsgGraphState) MaxPayloadLength(pver uint32) uint32 {
	return msg.GS.MaxPayloadLength()
}

func NewMsgGraphState(gs *dag.GraphState) *MsgGraphState {
	return &MsgGraphState{
		GS:gs,
	}
}