#ifndef HELLO_H_
#define HELLO_H_

#ifdef __cplusplus
extern "C"
{
#endif

    //返回可以听的牌，以逗号分隔
    extern const char *CheckTing(char *cards, int *tingCout);
    //胡牌检测
    //返回值说明 胡牌时返回 isWordReturn =0番型数值 =1返回字符。非胡牌,返回当前牌值。
    //参数说明
    // cards 玩家牌值(含露牌)
    // isMy (针对huCard)  =1时,huCard是玩家抓的牌; =0时,huCard是别家的牌。
    // huCard 单张牌值 (刚得到的牌值)
    // gang    是指上一手牌的操作为杠牌  别家开杠自己胡此牌(抢杠胡)也算。 
    // quanfeng 当前的风圈
    // menfeng 玩家的门风
    // haidi 牌墙中最后一张牌的牌值
    // code *code=1时,表示胡牌； *code=2表示自摸； =0,不能胡牌。
    // fanCout *fanCount 表示总计番数
    // isWordReturn =0番型数值 =1返回字符。非胡牌,返回当前牌值。
//    extern const char *CanHu(char *cards, int isMy, int huCard, int gang, int quanfeng, int menfeng, int juezhang, int haidi, int *code,int* fanCout,int isWordReturn);
    extern const char *CanHu(char *cards, int *code,int* fanCout,int isWordReturn);

#ifdef __cplusplus
}
#endif

#endif // HELLO_H_
