//build this program using below command 
//g++ client.cpp ./golibServer.so -o client 

#include "maingo.h"

#include <iostream>
#include <cstring>
#include <vector>
#include <string>
#include <iostream>



void SayHelloToCoder(std::string coder){
  //sends a string to the server, and the server "return" back a concatenated string 
   char *cptr = new char[coder.length()+1];
   std::strcpy (cptr, coder.c_str());

   cptr = GoSayHello(cptr);
   std::cout << cptr << std::endl;
   delete[] cptr;
}


void JoinStrings() {
  //sends two string to the server, server saves concatenated string in the third parameter.
   std::string s_in {"Let's rock "};
   std::vector<char> v_in {'G', 'o', '-', 'l', 'a', 'n', 'g', '!'};
   std::vector<char> v_out(20);

   GoString go_s_in{&s_in[0], static_cast<GoInt>(s_in.size())};
   GoSlice go_v_in{
        v_in.data(),
        static_cast<GoInt>(v_in.size()),
        static_cast<GoInt>(v_in.size()),
   };
   GoSlice go_v_out{
       v_out.data(),
       static_cast<GoInt>(v_out.size()),
       static_cast<GoInt>(v_out.size()),
   };

   GoConcatenate(go_s_in, go_v_in, go_v_out);//server will use go_v_out to update and retun value

  for(auto& c : v_out) {
        std::cout << c;
  }
  
  std::cout<<std::endl;
}

int main()
{
  SayHelloToCoder("Renjith");
  JoinStrings();

  std::string s_inFile {"/home/labuser/MyStuff/Readme"};
  GoString go_s_inFile{&s_inFile[0], static_cast<GoInt>(s_inFile.size())};
  GoPrintFileContent(go_s_inFile); //server will read and print the file content

  return 0;
}
