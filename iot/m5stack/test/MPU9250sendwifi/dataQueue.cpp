#include "dataQueue.h"



void dataQueue::enque(float x){

  Q[Qtail] = x;
if (Qtail == Qlenth ){
  Qtail =0;
  Qflag = true;
}
else{
  Qtail++;
}
  
}


bool  dataQueue::checkAwake(){
  if(Qflag == false){//qflagが立ってなかったら、qtailまでを判定
    int overNum = 0;//しきい値を超えた数
    for(int i=0; i<= Qtail;i++){
      if( - threshold > Q[i] || threshold < Q[i]){
        overNum++;
      } 
    }
    if(overNum >=thresholdAwakeNum){
        Qtail =0;
        Qflag = false;
      return true;
    }else{
    return false;

    }

  }else if(Qflag == true){//qflagが立ってたら、qlengthで判定
    int overNum = 0;//しきい値を超えた数
    for(int i=0; i<= Qlenth;i++){
      if( - threshold > Q[i] || threshold < Q[i]){
        overNum++;
      }
    }
    if(overNum >=thresholdAwakeNum){
        Qtail =0;
        Qflag = false;
      return true;
    }else{
    return false;

    }
    
  }
  
}

bool dataQueue::checkStable(){
  if(Qflag==false){//qflagが立ってなかったら、qtailまでを判定
    int overNum = 0;//しきい値を超えた数
    for(int i=0; i<= Qtail;i++){
      if(-threshold <Q[i] && threshold > Q[i] ){
        overNum++;
      }
    }
    if(overNum >=thresholdStableNum){
      Qtail =0;
      Qflag = false;
      return true;
    }else{
    return false;

    }

  }else if(Qflag == true){//qflagが立ってたら、qlengthで判定
    int overNum = 0;//しきい値を超えた数
    for(int i=0; i<= Qlenth;i++){
      if(-threshold <Q[i] && threshold > Q[i]){
        overNum++;
      }
    }
    if(overNum >=thresholdStableNum){
       Qtail =0;
       Qflag = false;
      return true;
    }else{
    return false;

    }
    
  }
  
}

