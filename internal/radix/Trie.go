// automatically generated by the FlatBuffers compiler, do not modify

package radix

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Trie struct {
	_tab flatbuffers.Table
}

func GetRootAsTrie(buf []byte, offset flatbuffers.UOffsetT) *Trie {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Trie{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Trie) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Trie) Nodes(obj *Node) *Node {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Node)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func TrieStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func TrieAddNodes(builder *flatbuffers.Builder, nodes flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(nodes), 0)
}
func TrieEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
