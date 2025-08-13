#include <stack>
#include <vector>
#include <string>
#include "free.h"
#include "matris.h"
#include "data.h"

using namespace std;

void handle_text(string text, Data &data);
bool handle_pr(string s);

Free::Free(Data &data){
	this->data = data;
}

void Free::loop(){
	bool over = false;
	
	cout << "You can use '!help' command!\n";
	while(!over){
		string command;
		
		cout << "\n$: ";
		getline(cin >> ws, command);
		cout << "\n";
		
		if(command == "!quit" || command == "!q"){
			over = true;
		}
		else if(command == "!help" || command == "!h"){
			this->help();
		}
		
		handle_text(command, data);
	}
}

void Free::help(){
	cout << "Kurallar\n";
	cout << "- Her Islem Parantez Ile Sarilmak Zorunda.\n";
	cout << "- Matris Islem Kurallarýna Dikkat Edilmeli.\n";
	cout << "- Burada Sadece Onceden Varolan Matrisleri Kullanabilirsiniz.\n";
	cout << "- Her Id Bir Harf Yerine Gecer. ( a:1 )\n";
	cout << "- Ex: ((a+b) * c) * d => 1,2,3 uncu Matrisleri Toplar 4. ile Carpar.\n";
	cout << "->!quit to exit!\n";
}

void handle_text(string text, Data &data){
	if(!handle_pr(text)){
		cout << "Parantez Kullanimi Hatali.";
		return;
	}
	
	vector<Matris> matrisler;
	
	// Get Variables
	for(int i = 0; i < text.length(); i++){
		if(text[i] > 96 && text[i] < 123){
			Matris x = data.byId(text[i]);
			matrisler.push_back(x);
		}
	}
}

bool handle_pr(string s){
	stack<char> prt;
	
	int size = s.length();

	int i = 0;
	while(i < size){
		if(s[i] == '{' || s[i] == '(' || s[i] == '['){
			prt.push(s[i]);
		}
		else{
			if(!prt.empty()){
				char top = prt.top();				
				if((int) top + 2 == (int) s[i] && s[i] != ')'){
					prt.pop();
				}
				else if((int) top + 1 == (int) s[i]){
					prt.pop();
				}
			}
		}
		i++;
	}
	
	if(prt.empty()){
		return true;
	}
	
	return false;
}

