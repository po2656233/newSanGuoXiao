#ifndef __CONSOLE_CONSOLE_H__
#define __CONSOLE_CONSOLE_H__

#include "mjpack.h"
#include "mjtile.h"
#include <cstdio>
#include <string>

namespace mahjong {

const std::string PackToEmojiString(const Pack &p);
const std::string TileToEmojiString(const Tile &p);

} // namespace mahjong

#endif
