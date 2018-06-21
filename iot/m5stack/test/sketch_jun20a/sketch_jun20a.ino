#include <Arduino.h>
//#include <WiFi.h>
#include <WiFiMulti.h>
#include <HTTPClient.h>
#include <M5Stack.h>
 
WiFiMulti wifi;
 
#define WIFI_SSID "yanaka"
#define WIFI_PASS "honma123"
const char *json = "{\"text\":\"hello\",\"num\":30}";
 

   
HTTPClient http;
 
void setup() {
 
  // put your setup code here, to run once:
  wifi.addAP(WIFI_SSID, WIFI_PASS);
  Serial.begin(115200);
  M5.begin();
  M5.Lcd.println("send slack");
   
  while (wifi.run() != WL_CONNECTED) {
    delay(500);
    M5.Lcd.printf(".");
  }
  M5.Lcd.println("wifi connect ok");
 
  // post slack
  M5.Lcd.println("connect slack.com");
  http.begin( "http://192.168.43.172:3000/" );
  http.addHeader("Content-Type", "application/json" );
  int a = http.POST((uint8_t*)json, strlen(json));
  M5.Lcd.println("post");
    M5.Lcd.println(a);

}
 
 
 
void loop() {
  // put your main code here, to run repeatedly:
 
}
