package chat

import (
	"strings"

	uv "github.com/charmbracelet/ultraviolet"
)

type Item interface {
	ID() string
	Render(width int) string
	Height(width int) int
}

type VirtualList struct {
	width, height int
	items         []Item
	gap           int

	offsetIdx   int
	offsetLine  int
	reverse     bool
	pausedAnims map[int]struct{}
}

type renderedItem struct {
	content string
	height  int
}

func NewVirtualList(items ...Item) *VirtualList {
	l := &VirtualList{
		items:       items,
		pausedAnims: make(map[int]struct{}),
	}
	return l
}

func (l *VirtualList) SetItems(items []Item) {
	if len(l.items) == 0 {
		l.items = items
		l.offsetIdx = 0
		l.offsetLine = 0
		return
	}

	l.items = items

	if l.offsetIdx >= len(l.items) {
		l.offsetIdx = max(0, len(l.items)-1)
		l.offsetLine = 0
	}
}

func (l *VirtualList) AppendItems(items ...Item) {
	if len(items) == 0 {
		return
	}
	l.items = append(l.items, items...)
}

func (l *VirtualList) UpdateItem(idx int, item Item) {
	if idx < 0 || idx >= len(l.items) {
		return
	}
	l.items[idx] = item
}

func (l *VirtualList) SetSize(width, height int) {
	l.width = width
	l.height = height
}

func (l *VirtualList) SetGap(gap int) {
	l.gap = gap
}

func (l *VirtualList) Gap() int {
	return l.gap
}

func (l *VirtualList) Width() int {
	return l.width
}

func (l *VirtualList) Height() int {
	return l.height
}

func (l *VirtualList) Len() int {
	return len(l.items)
}

func (l *VirtualList) SetReverse(reverse bool) {
	l.reverse = reverse
}

func (l *VirtualList) AtBottom() bool {
	if len(l.items) == 0 {
		return true
	}
	var remaining int
	for idx := l.offsetIdx; idx < len(l.items); idx++ {
		item := l.getItem(idx)
		if idx == l.offsetIdx {
			remaining += item.height - l.offsetLine
		} else {
			remaining += item.height
			if l.gap > 0 {
				remaining += l.gap
			}
		}
	}
	return remaining <= l.height
}

func (l *VirtualList) lastOffsetItem() (int, int, int) {
	var totalHeight int
	var idx int
	for idx = len(l.items) - 1; idx >= 0; idx-- {
		item := l.getItem(idx)
		itemHeight := item.height
		if l.gap > 0 && idx < len(l.items)-1 {
			itemHeight += l.gap
		}
		totalHeight += itemHeight
		if totalHeight > l.height {
			break
		}
	}
	lineOffset := max(totalHeight-l.height, 0)
	idx = max(idx, 0)
	return idx, lineOffset, totalHeight
}

func (l *VirtualList) getItem(idx int) renderedItem {
	if idx < 0 || idx >= len(l.items) {
		return renderedItem{}
	}
	item := l.items[idx]
	content := item.Render(l.width)
	content = strings.TrimRight(content, "\n")
	height := strings.Count(content, "\n") + 1
	return renderedItem{content: content, height: height}
}

func (l *VirtualList) ScrollToIndex(index int) {
	if index < 0 {
		index = 0
	}
	if index >= len(l.items) {
		index = len(l.items) - 1
	}
	if index < 0 {
		index = 0
	}
	l.offsetIdx = index
	l.offsetLine = 0
}

func (l *VirtualList) ScrollToTop() {
	l.offsetIdx = 0
	l.offsetLine = 0
}

func (l *VirtualList) ScrollToBottom() {
	if len(l.items) == 0 {
		return
	}
	lastOffsetIdx, lastOffsetLine, _ := l.lastOffsetItem()
	l.offsetIdx = lastOffsetIdx
	l.offsetLine = lastOffsetLine
}

func (l *VirtualList) ScrollBy(lines int) {
	if len(l.items) == 0 || lines == 0 {
		return
	}

	if l.reverse {
		lines = -lines
	}

	if lines > 0 {
		if l.AtBottom() {
			return
		}
		l.offsetLine += lines
		currentItem := l.getItem(l.offsetIdx)
		for l.offsetLine >= currentItem.height {
			l.offsetLine -= currentItem.height
			if l.gap > 0 {
				l.offsetLine = max(0, l.offsetLine-l.gap)
			}
			l.offsetIdx++
			if l.offsetIdx > len(l.items)-1 {
				l.ScrollToBottom()
				return
			}
			currentItem = l.getItem(l.offsetIdx)
		}
		lastOffsetIdx, lastOffsetLine, _ := l.lastOffsetItem()
		if l.offsetIdx > lastOffsetIdx || (l.offsetIdx == lastOffsetIdx && l.offsetLine > lastOffsetLine) {
			l.offsetIdx = lastOffsetIdx
			l.offsetLine = lastOffsetLine
		}
	} else if lines < 0 {
		if l.offsetIdx <= 0 && l.offsetLine <= 0 {
			return
		}
		l.offsetLine += lines
		for l.offsetLine < 0 {
			l.offsetIdx--
			if l.offsetIdx < 0 {
				l.offsetIdx = 0
				l.offsetLine = 0
				break
			}
			prevItem := l.getItem(l.offsetIdx)
			totalHeight := prevItem.height
			if l.gap > 0 {
				totalHeight += l.gap
			}
			l.offsetLine += totalHeight
		}
	}
}

func (l *VirtualList) PageDown() {
	l.ScrollBy(l.height)
}

func (l *VirtualList) PageUp() {
	l.ScrollBy(-l.height)
}

func (l *VirtualList) HalfPageUp() {
	l.ScrollBy(-l.height / 2)
}

func (l *VirtualList) HalfPageDown() {
	l.ScrollBy(l.height / 2)
}

func (l *VirtualList) ScrollUp() {
	l.ScrollBy(-1)
}

func (l *VirtualList) ScrollDown() {
	l.ScrollBy(1)
}

func (l *VirtualList) VisibleItemIndices() (startIdx, endIdx int) {
	if len(l.items) == 0 {
		return 0, 0
	}
	startIdx = l.offsetIdx
	currentIdx := startIdx
	visibleHeight := -l.offsetLine

	for currentIdx < len(l.items) {
		item := l.getItem(currentIdx)
		visibleHeight += item.height
		if l.gap > 0 {
			visibleHeight += l.gap
		}
		if visibleHeight >= l.height {
			break
		}
		currentIdx++
	}

	endIdx = currentIdx
	if endIdx >= len(l.items) {
		endIdx = len(l.items) - 1
	}
	if endIdx < 0 {
		endIdx = 0
	}

	return startIdx, endIdx
}

func (l *VirtualList) Render() string {
	if len(l.items) == 0 {
		return ""
	}

	var lines []string
	currentIdx := l.offsetIdx
	currentOffset := l.offsetLine
	linesNeeded := l.height

	for linesNeeded > 0 && currentIdx < len(l.items) {
		item := l.getItem(currentIdx)
		itemLines := strings.Split(item.content, "\n")
		itemHeight := len(itemLines)

		if currentOffset < 0 {
			negativeOffset := -currentOffset
			if negativeOffset >= itemHeight {
				if l.gap > 0 {
					gapScrolled := negativeOffset - itemHeight
					gapRemaining := l.gap - gapScrolled
					for i := 0; i < gapRemaining && linesNeeded > 0; i++ {
						lines = append(lines, "")
						linesNeeded--
					}
				}
			} else {
				visibleLines := itemHeight - negativeOffset
				end := min(visibleLines, len(itemLines))
				lines = append(lines, itemLines[:end]...)
				linesNeeded -= end
			}
		} else if currentOffset == 0 {
			lines = append(lines, itemLines...)
			linesNeeded -= itemHeight
			if l.gap > 0 && linesNeeded > 0 {
				lines = append(lines, "")
				linesNeeded--
			}
		} else if currentOffset < itemHeight {
			lines = append(lines, itemLines[currentOffset:]...)
			linesNeeded -= len(itemLines) - currentOffset
			if l.gap > 0 && linesNeeded > 0 {
				lines = append(lines, "")
				linesNeeded--
			}
		} else {
			if l.gap > 0 && linesNeeded > 0 {
				gapRemaining := l.gap
				for i := 0; i < gapRemaining && linesNeeded > 0; i++ {
					lines = append(lines, "")
					linesNeeded--
				}
			}
		}

		currentIdx++
		currentOffset = 0
	}

	if len(lines) > l.height {
		lines = lines[:l.height]
	}

	if l.reverse {
		for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
			lines[i], lines[j] = lines[j], lines[i]
		}
	}

	return strings.Join(lines, "\n")
}

func (l *VirtualList) Draw(scr uv.Screen, area uv.Rectangle) {
	content := l.Render()
	uv.NewStyledString(content).Draw(scr, area)
}
