#include <stdio.h>
#include "mjwrapper.h"


//返回可以听的牌，以逗号分隔
const char *CheckTing(char *cards, int *tingCout)
{
    return call_MJ_CheckTing(cards, tingCout);
}

//是否可以胡牌
const char *CanHu(char *cards, int *code,int* fanCout,int isWordReturn)
{
   return call_MJ_CanHu(cards,  code,fanCout,isWordReturn);
}
