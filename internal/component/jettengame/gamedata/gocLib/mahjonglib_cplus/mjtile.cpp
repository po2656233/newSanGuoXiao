#include "mjtile.h"

namespace mahjong {

/*
牌的Unicode顺序：
🀀🀁🀂🀃🀄🀅🀆🀇🀈🀉🀊🀋🀌🀍🀎🀏🀐🀑🀒🀓🀔🀕🀖🀗🀘🀙🀚🀛🀜🀝🀞🀟🀠🀡🀢🀣🀤🀥🀦🀧🀨🀩🀪🀫
🀀 🀁 🀂 🀃 🀄 🀅 🀆 🀇 🀈 🀉 🀊 🀋 🀌 🀍 🀎 🀏 🀐 🀑 🀒 🀓 🀔 🀕 🀖 🀗 🀘 🀙 🀚 🀛 🀜 🀝 🀞 🀟 🀠 🀡 🀢 🀣 🀤 🀥 🀦 🀧 🀨 🀩 🀪 🀫
*/

const char *TILES_UTF8[] = {
    "",
    "🀇", "🀈", "🀉", "🀊", "🀋", "🀌", "🀍", "🀎", "🀏",
    "🀐", "🀑", "🀒", "🀓", "🀔", "🀕", "🀖", "🀗", "🀘",
    "🀙", "🀚", "🀛", "🀜", "🀝", "🀞", "🀟", "🀠", "🀡",
    "🀀", "🀁", "🀂", "🀃",
    "🀄", "🀅", "🀆",
    "🀢", "🀣", "🀤", "🀥", "🀦", "🀧", "🀨", "🀩",
    "🀪", "🀫"};

const unsigned TILES_SUIT[] = {
    SUIT_INVALID,
    SUIT_WAN, SUIT_WAN, SUIT_WAN, SUIT_WAN, SUIT_WAN, SUIT_WAN, SUIT_WAN, SUIT_WAN, SUIT_WAN,
    SUIT_TIAO, SUIT_TIAO, SUIT_TIAO, SUIT_TIAO, SUIT_TIAO, SUIT_TIAO, SUIT_TIAO, SUIT_TIAO, SUIT_TIAO,
    SUIT_BING, SUIT_BING, SUIT_BING, SUIT_BING, SUIT_BING, SUIT_BING, SUIT_BING, SUIT_BING, SUIT_BING,
    SUIT_FENG, SUIT_FENG, SUIT_FENG, SUIT_FENG,
    SUIT_JIAN, SUIT_JIAN, SUIT_JIAN,
    SUIT_HUA, SUIT_HUA, SUIT_HUA, SUIT_HUA, SUIT_HUA, SUIT_HUA, SUIT_HUA, SUIT_HUA,
    SUIT_INVALID, SUIT_INVALID};

const unsigned TILES_RANK[] = {
    RANK_INVALID,
    RANK_1, RANK_2, RANK_3, RANK_4, RANK_5, RANK_6, RANK_7, RANK_8, RANK_9,
    RANK_1, RANK_2, RANK_3, RANK_4, RANK_5, RANK_6, RANK_7, RANK_8, RANK_9,
    RANK_1, RANK_2, RANK_3, RANK_4, RANK_5, RANK_6, RANK_7, RANK_8, RANK_9,
    RANK_INVALID, RANK_INVALID, RANK_INVALID, RANK_INVALID,
    RANK_INVALID, RANK_INVALID, RANK_INVALID,
    RANK_INVALID, RANK_INVALID, RANK_INVALID, RANK_INVALID, RANK_INVALID, RANK_INVALID, RANK_INVALID, RANK_INVALID,
    RANK_INVALID, RANK_INVALID};

const char TILES_SUIT_CHAR[] = {
    TILE_CHAR_INVALID,
    TILE_CHAR_WAN, TILE_CHAR_WAN, TILE_CHAR_WAN, TILE_CHAR_WAN, TILE_CHAR_WAN, TILE_CHAR_WAN, TILE_CHAR_WAN, TILE_CHAR_WAN, TILE_CHAR_WAN,
    TILE_CHAR_TIAO, TILE_CHAR_TIAO, TILE_CHAR_TIAO, TILE_CHAR_TIAO, TILE_CHAR_TIAO, TILE_CHAR_TIAO, TILE_CHAR_TIAO, TILE_CHAR_TIAO, TILE_CHAR_TIAO,
    TILE_CHAR_BING, TILE_CHAR_BING, TILE_CHAR_BING, TILE_CHAR_BING, TILE_CHAR_BING, TILE_CHAR_BING, TILE_CHAR_BING, TILE_CHAR_BING, TILE_CHAR_BING,
    TILE_CHAR_E, TILE_CHAR_S, TILE_CHAR_W, TILE_CHAR_N,
    TILE_CHAR_C, TILE_CHAR_F, TILE_CHAR_P,
    TILE_CHAR_MEI, TILE_CHAR_LAN, TILE_CHAR_ZHU, TILE_CHAR_JU, TILE_CHAR_CHU, TILE_CHAR_XIA, TILE_CHAR_QIU, TILE_CHAR_DONG,
    TILE_CHAR_INVALID, TILE_CHAR_INVALID};

} // namespace mahjong