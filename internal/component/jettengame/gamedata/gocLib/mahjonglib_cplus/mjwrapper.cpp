#include "mjhandtiles.h"
#include "mjwrapper.h"
#include "mjfan.h"
#include "mjtile.h"
#include "mjprint.h"
#include "mjconsole.h"
 #include <cstring>

using namespace std;

// 检测听牌
const char *call_MJ_CheckTing(char *cards, int *tingCout)
{
  mahjong::Handtiles ht;
  ht.StringToHandtiles(cards);
  // 计算听牌
  mahjong::Fan fan;
  std::vector<mahjong::Tile> ting = fan.CalcTing(ht);
  std::string ting_string("");
  int nCout = 0;
  for (mahjong::Tile t : ting)
  {
    nCout++;
    //mahjong::StdPrintTile(t);
    ting_string += t.UTF8();
    ting_string += ",";
  }
  *tingCout = nCout;
  char *p = new char[256];
  memset(p, 0, 256);
  strncpy(p, ting_string.c_str(),ting_string.length());
  return p;
}

//胡牌检测 除了胡牌 其他牌值默认为0
const char *call_MJ_CanHu(char *cards, int *code, int *fanCout, int isWordReturn)
{
  // 判断是否铳和
  mahjong::Handtiles ht;
  ht.StringToHandtiles(cards);

  mahjong::Fan fan;
  int ret = fan.JudgeHu(ht);
  char *p = new char[256];

  memset(p, 0, 256);
  *code = ret; // 胡牌

  if (ret == 1)
  {
    fan.CountFan(ht);
    *fanCout = fan.tot_fan_res; // 计番

    char szFan[64];
    std::string strFan("");
    int index = 0;
    for (int i = 1; i < mahjong::FAN_SIZE; i++)
    { // 输出所有的番
      for (size_t j = 0; j < fan.fan_table_res[i].size(); j++)
      {
        if (isWordReturn == 0)
        {
          p[index++] = i;
          continue;
        }
        memset(szFan, 0, sizeof(szFan));
        sprintf(szFan,  "%s %d番", mahjong::FAN_NAME[i], mahjong::FAN_SCORE[i]);
        strFan += szFan;
        std::string pack_string;
        for (auto pid : fan.fan_table_res[i][j])
        { // 获取该番种具体的组合方式
          pack_string += " " + mahjong::PackToEmojiString(fan.fan_packs_res[pid]);
        }
        strFan += pack_string + "\n";

        // printf("%s\n", pack_string.c_str());
      }
    }
    if (isWordReturn == 1)
    {
      strncpy(p, strFan.c_str(),strFan.length());
    }
  }
  else
  {
    string s = ht.HandtilesToString();
    strncpy(p,  s.c_str(),s.length());
  }

  return p;
}

// int call_MJ_CanHu(void *p, int tile)
// {
//   mahjong::Handtiles *ht = static_cast<mahjong::Handtiles *>(p);
//   ht->SetTile(mahjong::Tile(tile));
//   mahjong::Fan fan;
//   int code = fan.JudgeHu(*ht);
//   printf("%s 是否和牌：%d\n", ht->HandtilesToString().c_str(), code);
//   ht->DiscardTile(); //

//   return code;
// }

// int call_MJ_JiFan(void *p, int tile)
// {
//   mahjong::Handtiles *ht = static_cast<mahjong::Handtiles *>(p);
//   ht->DrawTile(mahjong::Tile(tile));                     // 摸牌
//   printf("手牌：%s\n", ht->HandtilesToString().c_str()); //
//   mahjong::Fan fan;
//   fan.CountFan(*ht); // 计番

//   int jiFan = 0;
//   printf("总番数：%d\n", fan.tot_fan_res); // 总番数
//   for (int i = 1; i < mahjong::FAN_SIZE; i++)
//   { // 输出所有的番
//     for (size_t j = 0; j < fan.fan_table_res[i].size(); j++)
//     {
//       printf("%s %d番", mahjong::FAN_NAME[i], mahjong::FAN_SCORE[i]);
//       jiFan += mahjong::FAN_SCORE[i];
//       std::string pack_string;
//       for (auto pid : fan.fan_table_res[i][j])
//       { // 获取该番种具体的组合方式
//         pack_string += " " + mahjong::PackToEmojiString(fan.fan_packs_res[pid]);
//       }
//       printf("%s\n", pack_string.c_str());
//     }
//   }
//   return jiFan;
// }
