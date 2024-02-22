package po

type token int

const (
	eof                    token = iota // "eof"
	whitespace                          // "ws"
	failure                             // "failure"
	commentTranslator                   // "#"
	commentExtracted                    // "#."
	commentReference                    // "#:"
	commentFlags                        // "#,"
	commentPrevContext                  // "#| msgctxt"
	commentPrevContextLine              // "#| \"prev-ctx\""
	commentPrevMsgID                    // "#| msgid"
	commentPrevMsgIDLine                // "#| \"prev-msg\""
	commentPrevUnknown                  // "#| \"unknown\""
	msgContext                          // "msgctxt"
	msgContextLine                      // "\"more context\""
	msgID                               // "msgid"
	msgIDLine                           // "\"singular\""
	msgIDPlural                         // "msgid_plural"
	msgIDPluralLine                     // "\"plural\""
	msgStr                              // "msgstr"
	msgStrLine                          // "\"singular translation\""
	msgStrPlural                        // "msgstr[?]"
	msgStrPluralLine                    // "\"plural translation\""
	none                                // internal state
)
