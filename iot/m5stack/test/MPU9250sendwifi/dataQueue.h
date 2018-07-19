#ifndef DATAQUEUE_H_INCLUDE
#define DATAQUEUE_H_INCLUDE

class dataQueue {
  public:

    float threshold = 50;//しきい値
    int thresholdAwakeNum = 7;//しきい値を何回超えたか。
    int thresholdStableNum = 20;//しきい値を何回超えたか。
    int Qlenth = 30;  //長さ
    void enque(float x);
    bool getQflag();//GetFlag
    bool checkAwake();// 起動したかのチェック
    bool checkStable();//安定したかのチェック
  private:
    bool Qflag = 0;
    float Q[100];
    int Qtail =  0;//おわり
};

#endif
