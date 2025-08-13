#include <iostream>
#include <cmath>
#include <vector>
#include "data.h"
#include "matris.h"
#include "window.h"
#include "free.h"

using namespace std;

int* getSize(string &text);
vector<Matris> getMatris(Data &data);
bool canAdd(Matris x, Matris y);
bool canMultiply(Matris x, Matris y);
bool canDet(Matris x);
int validId(int index, Data &data);

void Window::loop(string lang){
	bool over = false;
	string help = lang == "eu" ? "You can use '!help' command!\n":"Yardim istemek icin '!yardim'' komutunu kullanabilirsin.";
	
	Data data;
	
	// Static Method
	Matris::takeFromList(data);
	
	
	
	cout << help;
	while(!over){
		Free free(data);
		string command;
		
		cout << "\n>: ";
		getline(cin >> ws, command);
		cout << "\n\n";
		
		if(command == "!quit" || command == "!q"){
			over = true;
		}
		else if(command == "!help" || command == "!yardim"){
			this->help(lang);
		}
		else if(command == "!create" || command == "!c" || command == "!yap"){
			this->create(data);
		}
		else if(command == "!list" || command == "!liste"){
			data.listData();
		}
		else if(command == "!add" || command == "!topla"){
			this->add(data);
		}
		else if(command == "!multiply" || command == "!carp"){
			this->multiply(data);
		}
		else if(command == "!det"){
			this->det(data);
		}
		/*
		else if(command == "!free"){			
			free.loop();
		}
		*/
	}
}

void Window::help(string lang){
	if(lang == "tr"){
		cout << "- Matris olusturmak icin !yap yaz.\n";
		cout << "- Veritabanindaki matrisleri gormek icin !liste yaz.\n";
		cout << "- Matrisleri toplamak icin !topla yaz.\n";
		cout << "- Matrisleri carpmak icin !carp yaz.\n";
		cout << "- Matrislerin determinantýný almak icin !det yaz.\n";
		//cout << "- !free to calculate easily.\n";
		cout << "-> Cikmak icin !quit / !q yaz.\n";
	}
	else if(lang == "eu"){
		cout << "- !create / !c to create a matris.\n";
		cout << "- !list to list all matris in database.\n";
		cout << "- !add to add matrises.\n";
		cout << "- !multiply to add matrises.\n";
		cout << "- !det to determinant matris.\n";
		//cout << "- !free to calculate easily.\n";
		cout << "-> !quit / !q to exit.\n";
	}
}

void Window::multiply(Data &data){
	int dataSize = data.getSize();
	
	if(dataSize < 1){
		cout << "Not Enough Matrices In Database.";
		return;
	}
	vector<Matris> matris = getMatris(data);
	
	Matris x = matris[0], y = matris[1];
	
	if(x.getId() == -1 || y.getId() == -1){
		cout << "There is no such id.";
		return;
	}
	
	if(!canMultiply(x, y)){
		cout << "For matrix multiplication, the row of the first matrix and the column of the second must be equal.";
		return;
	}
	
	cout << "\nResult\n";
	(x*y).printMatris();
}

void Window::add(Data &data){
	int dataSize = data.getSize();
	
	if(dataSize < 1){
		cout << "Not Enough Matrices In Database.";
		return;
	}
	vector<Matris> matris = getMatris(data);
	
	Matris x = matris[0], y = matris[1];
	
	if(x.getId() == -1 || y.getId() == -1){
		cout << "There is no such id.";
		return;
	}
	
	if(!canAdd(x, y)){
		cout << "rows or columns of matrices do not match.";
		return;
	}
	
	cout << "Result\n";
	(x + y).printMatris();
}

void Window::create(Data &data){
	int *size = new int[2];
	string size_text;
	
	cout << "Format => 3x3 / 5x3\n";
	cout << "Enter the Size: ";
	cin >> size_text;
	
	size = getSize(size_text);
	//cout << size[0] << " " << size[1] << endl;
	
	// Creating a matris.
	vector<double> value_row;
	arr2d values;
	double num;
	
	for(int i = 0; i < size[0]; i++){
		for(int j = 0; j < size[1]; j++){
			cout << "Enter the [ " << i << " ] [ " << j << " ] Value: " ;
			cin >> num;
			value_row.push_back(num);
		}
		
		values.push_back(value_row);
		value_row.clear();
	}
	Matris x(values);
	values.clear();
	data.addData(x);
	
}

void Window::det(Data &data){
	int dataSize = data.getSize();
	
	if(dataSize < 1){
		cout << "Not Enough Matrices In Database.";
		return;
	}
	int id;
	
	cout << "Choose From ID \n";
	data.listData();
	cout << "\n\n";
	
	cout << "ID: ";
	cin >> id;
	
	Matris x = data.chooseMatris(id);
	
	if(x.getId() == -1){
		cout << "There is no such id.";
		return;
	}
	
	if(!canDet(x)){
		cout << "The number of rows and the number of columns of the matrix must be equal.";
		return;
	}
	
	x.printMatris();
	cout << "Determinant: " << x.determinant();
}

// UTILS

bool canDet(Matris x){
	int* xSize = new int[2];
	
	xSize = x.getSize();
	
	if(xSize[0] != xSize[1]){
		return false;
	}
	
	return true;
}

bool canMultiply(Matris x, Matris y){
	int* xSize = new int[2];
	int* ySize = new int[2];
	
	xSize = x.getSize();
	ySize = y.getSize();
	
	if(xSize[1] != ySize[0]){
		return false;
	}
	
	return true;
}

bool canAdd(Matris x, Matris y){
	int* xSize = new int[2];
	int* ySize = new int[2];
	
	xSize = x.getSize();
	ySize = y.getSize();
	
	if(xSize[0] != ySize[0] || xSize[1] != ySize[1]){
		return false;
	}
	
	return true;
}

vector<Matris> getMatris(Data &data){
	Matris Error(-1, 1, 1, true);
	vector<Matris> matris;
	
	int first, second;
	
	cout << "Choose From ID \n";
	data.listData();
	cout << "\n\n";
	
	cout << "1. Matris: ";
	cin >> first;
	
	cout << "2. Matris: ";
	cin >> second;
	
	if(validId(first, data) == false && validId(second, data) == false){
		cout << "There is no such id.";
		matris.push_back(Error);
		return matris;
	}
	
	Matris x = data.chooseMatris(first);
	Matris y = data.chooseMatris(second);
	
	matris.push_back(x);
	matris.push_back(y);
	
	return matris;
}

int validId(int index, Data &data){
	int* validIds = new int[data.getSize()];
	validIds = data.validIds();
	
	for(int i = 0; i < data.getSize(); i++){
		if(index == validIds[i]){
			return true;
		}
	}
	
	return false;
}

int* getSize(string &text){
	int* size = new int[2];
	
	// For Error Handling
	size[0] = -1, size[1] = -1;
	
	int index;
	
	string row = "";
	string col = "";
	
	// Getting the number part of the text
	while(true){
		for(int i = 0; i < text.length(); i++){
			if(text[i] == 'x'){
				index = i;
				break;
			}
				row += text[i];
		}
		
		for(int i = index + 1; i < text.length(); i++){
			col += text[i];
		}
		
		break;
	}
	
	// Converting row and col to int
	for(int i = 0; i < row.length(); i++){
		size[0] = (row[i] - '0') * pow(10, row.length() - i - 1);
	}
	
	for(int j = 0; j < col.length(); j++){
		size[1] = (col[j] - '0') * pow(10, col.length() - j - 1);
	}
	
	return size;
}

