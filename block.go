package minecraft

import (
	"fmt"

	"github.com/MJKWoolnough/equaler"
	"github.com/MJKWoolnough/minecraft/nbt"
)

// Tick is a type that represents a scheduled tick
type Tick struct {
	I, T, P int32
}

// Block is a type that represents the full information for a block, id, data,
// metadata and scheduled tick data
type Block struct {
	BlockID  uint16
	Data     uint8
	metadata nbt.Compound
	ticks    []Tick
}

// Equal is an implementation of the equaler.Equaler interface
func (b Block) Equal(e equaler.Equaler) bool {
	c, ok := e.(*Block)
	if !ok {
		if d, ok := e.(Block); ok {
			c = &d
		}
	}
	if c != nil {
		if b.BlockID == c.BlockID && b.Data == c.Data && len(b.metadata) == len(c.metadata) && len(b.ticks) == len(c.ticks) {
			for _, bT := range b.ticks {
				found := false
				for _, cT := range c.ticks {
					if bT.I == cT.I && bT.T == cT.T && bT.P == cT.P {
						found = true
						break
					}
				}
				if !found {
					return false
				}
			}
			for _, v := range b.metadata {
				name := v.Name()
				found := false
				for _, w := range c.metadata {
					if w.Name() == name {
						if !v.Equal(w) {
							return false
						}
						found = true
						break
					}
				}
				if !found {
					return false
				}
			}
			return true
		}
	}
	return false
}

// Opacity returns how much light is blocked by this block.
func (b *Block) Opacity() uint8 {
	if b.BlockID == 8 || b.BlockID == 9 {
		return 3
	}
	for i := 0; i < len(TransparentBlocks); i++ {
		if TransparentBlocks[i] == b.BlockID {
			return 1
		}
	}
	return 15
}

// Light returns how much light is generated by this block.
func (b *Block) Light() uint8 {
	if l, ok := LightBlocks[b.BlockID]; ok {
		return l
	}
	return 0
}

// IsLiquid returns true if the block id matches a liquid
func (b *Block) IsLiquid() bool {
	return b.BlockID == 8 || b.BlockID == 9 || b.BlockID == 10 || b.BlockID == 11
}

// HasMetadata returns true the the block contains extended metadata
func (b *Block) HasMetadata() bool {
	return len(b.metadata) > 0
}

// GetMetadata returns a copy of the metadata for this block, or nil is it has
// none
func (b *Block) GetMetadata() nbt.Compound {
	if b.metadata == nil {
		return nil
	}
	return *b.metadata.Copy().(*nbt.Compound)
}

// SetMetadata sets the blocks metadata to a copy of the given metadata
func (b *Block) SetMetadata(data nbt.Compound) {
	metadata := make(nbt.Compound, 0)
	for i := 0; i < len(data); i++ {
		name := data[i].Name()
		if name == "x" || name == "y" || name == "z" {
			continue
		} else if data[i].TagID() == nbt.TagEnd {
			break
		}
		metadata = append(metadata, data[i].Copy())
	}
	if len(metadata) > 0 {
		b.metadata = metadata
	} else {
		b.metadata = nil
	}
}

// HasTicks returns true if the block has any scheduled ticks
func (b *Block) HasTicks() bool {
	return len(b.ticks) > 0
}

// GetTicks returns all of the scheduled ticks for a block
func (b *Block) GetTicks() []Tick {
	t := make([]Tick, len(b.ticks))
	copy(t, b.ticks)
	return t
}

// AddTicks adds one or more scheduled ticks to the block
func (b *Block) AddTicks(t ...Tick) {
	b.ticks = append(b.ticks, t...)
}

// SetTicks sets the scheduled ticks for the block, replacing any existing ones
func (b *Block) SetTicks(t []Tick) {
	b.ticks = make([]Tick, len(t))
	copy(b.ticks, t)
}

func (b *Block) String() string {
	toRet := fmt.Sprintf("Block ID: %d\nData: %d\n", b.BlockID, b.Data)
	if b.metadata != nil && len(b.metadata) != 0 {
		toRet += "Metadata: " + b.metadata.String()
	}
	for n, tick := range b.ticks {
		toRet += fmt.Sprintf("	Tick: %d, i: %d, t: %d, p: %d\n", n+1, tick.I, tick.T, tick.P)
	}
	return toRet
}
