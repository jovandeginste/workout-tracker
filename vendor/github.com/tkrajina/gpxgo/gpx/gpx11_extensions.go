package gpx

import (
	"encoding/xml"
	"strings"
)

type ExtensionNode struct {
	XMLName xml.Name
	Attrs   []xml.Attr      `xml:",any,attr"`
	Data    string          `xml:",chardata"`
	Nodes   []ExtensionNode `xml:",any"`
}

func (n ExtensionNode) debugXMLChunk() []byte {
	byts, err := xml.MarshalIndent(n, "", "    ")
	if err != nil {
		return []byte("???")
	}
	return byts
}

func (n ExtensionNode) toTokens(prefix string) (tokens []xml.Token) {
	var attrs []xml.Attr
	for _, a := range n.Attrs {
		attrs = append(attrs, xml.Attr{Name: xml.Name{Local: prefix + a.Name.Local}, Value: a.Value})
	}

	start := xml.StartElement{Name: xml.Name{Local: prefix + n.XMLName.Local, Space: ""}, Attr: attrs}
	tokens = append(tokens, start)
	data := strings.TrimSpace(n.Data)
	if len(n.Nodes) > 0 {
		for _, node := range n.Nodes {
			tokens = append(tokens, node.toTokens(prefix)...)
		}
	} else if data != "" {
		tokens = append(tokens, xml.CharData(data))
	} else {
		return nil
	}
	tokens = append(tokens, xml.EndElement{start.Name})
	return
}

func (n *ExtensionNode) GetAttr(key string) (value string, found bool) {
	for i := range n.Attrs {
		if n.Attrs[i].Name.Local == key {
			value = n.Attrs[i].Value
			found = true
			return
		}
	}
	return
}

func (n *ExtensionNode) SetAttr(key, value string) {
	for i := range n.Attrs {
		if n.Attrs[i].Name.Local == key {
			n.Attrs[i].Value = value
			return
		}
	}
	n.Attrs = append(n.Attrs, xml.Attr{
		Name: xml.Name{
			Space: n.SpaceNameURL(),
			Local: key,
		},
		Value: value,
	})
}

func (n *ExtensionNode) GetNode(path0 string) (node *ExtensionNode, found bool) {
	for subn := range n.Nodes {
		if n.Nodes[subn].LocalName() == path0 {
			node = &n.Nodes[subn]
			found = true
			return
		}
	}
	return
}

func (n *ExtensionNode) GetOrCreateNode(path ...string) *ExtensionNode {
	if len(path) == 0 {
		return n
	}

	path0, rest := path[0], path[1:]

	subNode, found := n.GetNode(path0)
	if !found {
		n.Nodes = append(n.Nodes, ExtensionNode{
			XMLName: xml.Name{
				Space: n.XMLName.Space,
				Local: path0,
			},
			Attrs: nil,
		})
		subNode = &(n.Nodes[len(n.Nodes)-1])
	}

	return subNode.GetOrCreateNode(rest...)
}

func (n ExtensionNode) IsEmpty() bool {
	return len(n.Nodes) == 0 && len(n.Attrs) == 0 && len(n.Data) == 0
}
func (n ExtensionNode) LocalName() string    { return n.XMLName.Local }
func (n ExtensionNode) SpaceNameURL() string { return n.XMLName.Space }
func (n ExtensionNode) GetAttrOrEmpty(attr string) string {
	val, _ := n.GetAttr(attr)
	return val
}

type Extension struct {
	// XMLName xml.Name
	// Attrs   []xml.Attr `xml:",any,attr"`
	Nodes []ExtensionNode `xml:",any"`

	// Filled before deserializing:
	globalNsAttrs map[string]NamespaceAttribute
}

var _ xml.Marshaler = Extension{}

func (ex Extension) debugXMLChunk() []byte {
	byts, err := xml.MarshalIndent(ex, "", "    ")
	if err != nil {
		return []byte("???")
	}
	return byts
}

func (ex Extension) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(ex.Nodes) == 0 {
		return nil
	}

	start = xml.StartElement{Name: xml.Name{Local: start.Name.Local}, Attr: nil}
	tokens := []xml.Token{start}
	for _, node := range ex.Nodes {
		prefix := ""
		for _, v := range ex.globalNsAttrs {
			if node.SpaceNameURL() == v.Value || node.SpaceNameURL() == v.Name.Local {
				prefix = v.replacement
			}
		}
		tokens = append(tokens, node.toTokens(prefix)...)
	}

	tokens = append(tokens, xml.EndElement{Name: start.Name})

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	// flush to ensure tokens are written
	err := e.Flush()
	if err != nil {
		return err
	}

	return nil
}

type NamespaceURL string

const (
	// NoNamespace is used for extension nodes without namespace
	NoNamespace NamespaceURL = ""
	// AnyNamespace is an invalid namespace used for searching for nodes by name (regardless of namespace)
	AnyNamespace NamespaceURL = "-1"
)

func (ex *Extension) GetOrCreateNode(namespaceURL NamespaceURL, path ...string) *ExtensionNode {
	// TODO: Check is len(nodes) == 0
	var subNode *ExtensionNode
	for n := range ex.Nodes {
		if ex.Nodes[n].SpaceNameURL() == string(namespaceURL) && ex.Nodes[n].LocalName() == path[0] {
			subNode = &ex.Nodes[n]
			break
		}
	}
	if subNode == nil {
		ex.Nodes = append(ex.Nodes, ExtensionNode{
			XMLName: xml.Name{
				Space: string(namespaceURL),
				Local: path[0],
			},
		})
		subNode = &ex.Nodes[len(ex.Nodes)-1]
	}
	return subNode.GetOrCreateNode(path[1:]...)
}

func (ex *Extension) GetNode(namespaceURL NamespaceURL, path0 string) (node *ExtensionNode, found bool) {
	for subn := range ex.Nodes {
		if ex.Nodes[subn].LocalName() == path0 {
			if ex.Nodes[subn].SpaceNameURL() == string(namespaceURL) || namespaceURL == AnyNamespace {
				node = &ex.Nodes[subn]
				found = true
				return
			}
		}
	}
	return
}
