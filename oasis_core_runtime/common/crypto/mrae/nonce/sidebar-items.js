window.SIDEBAR_ITEMS = {"constant":[["NONCE_SIZE","Size of the nonce in bytes. Size of the nonce in bytes."],["TAG_SIZE","Size of tag portion of the nonce in bytes. These bytes will never update."]],"struct":[["Nonce","120 bit nonce with a 88 bit tag and 32 bit counter. If the counter exceeds 32 bits, then the nonce is no longer valid and must be refreshed with a new random nonce. It is expected that all 128 bits are given randomly. However, the last 32 counting bits may wrap around to ensure 2^32 counts may be used per nonce."]]};