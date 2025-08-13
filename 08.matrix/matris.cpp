#include <iostream>
#include <vector>
#include <cmath>
#include <fstream>
#include "matris.h"
#include "data.h"

Matris::Matris(arr2d Value){
	this->id = -1;
	this->value = Value;
}

Matris::Matris(int size){
	double arr[size][size];
	
	for(int i = 0; i < size; i++){
		for(int j = 0; j < size; j++){
			if(i == j)
				arr[i][j] = 1;
			else
				arr[i][j] = 0;
		}
	}
	
	vector<double> newValue;
	
	for(int i = 0; i < size; i++){
		for(int j = 0; j < size; j++){
			newValue.push_back(arr[i][j]);
		}
		this->value.push_back(newValue);
		newValue.clear();
	}
}

Matris::Matris(double number, int row, int col, bool error){
	double arr[row][col];
	
	for(int i = 0; i < row; i++){
		for(int j = 0; j < col; j++){
			arr[i][j] = number;
		}
	}
	
	vector<double> newValue;
	
	for(int i = 0; i < row; i++){
		for(int j = 0; j < col; j++){
			newValue.push_back(arr[i][j]);
		}
		this->value.push_back(newValue);
		newValue.clear();
	}
	
	if(error == true){
		this->id = -1;
	}
}

void Matris::takeFromList(Data &data){
	// ilk 2 input row col sonrasý deðerler
	ifstream input("MatrisList.txt");
	double value;
	
	
	while(input >> value){
		int row = value;
		
		input >> value;
		int col = value;
		
		vector<vector<double>> matris;
		vector<double> matris_row;
		
		for(int i = 0; i < row; i++){
			for(int j = 0; j < col; j++){
				input >> value;
				matris_row.push_back(value);
			}
			matris.push_back(matris_row);
			matris_row.clear();
		}
		
		Matris x(matris);
		
		data.addData(x);
	}
	
	return;
	
}

int* Matris::getSize(){
	int* arr = new int[2];
	
	arr[0] = this->value.size();
	arr[1] = this->value[0].size();
	
	return arr;
}

void Matris::printMatris(){
	int* size = new int[2];
	size = this->getSize();
	
	cout << " \n";
	for(int i = 0; i < size[0]; i++){
		cout << "[ ";
		for(int j = 0; j < size[1]; j++){
			cout << this->value[i][j] << " ";
		}
		cout << "]\n";
	}
}

void Matris::listMatris(){
	int* size = new int[2];
	size = this->getSize();
	
	cout << "\n" << this->id << ": ";
	for(int i = 0; i < size[0]; i++){
		cout << "[ ";
		for(int j = 0; j < size[1]; j++){
			cout << this->value[i][j] << " ";
		}
		cout << "]";
	}
}

Matris Matris::operator+(Matris &other){
	Matris Error(-1, 2, 2, true);
	int* aSize = new int[2];
	int* bSize = new int[2];
	
	aSize = this->getSize();
	bSize = other.getSize();
	
	Matris result(0, aSize[0], aSize[1]);
	
	if(aSize[0] != bSize[0] || aSize[1] != bSize[1]){
		cout << "Size's Do not Match. Error:1\n";
		return Error;
	}
		
	
	for(int i = 0; i < aSize[0]; i++){
		for(int j = 0; j < aSize[1]; j++){
			result.value[i][j] = this->value[i][j] + other.value[i][j];
		}
	}
	
	return result;
}

Matris Matris::operator*(Matris &other){
	Matris Error(-1, 2, 2, true);
	int* aSize = new int[2];
	int* bSize = new int[2];
	
	aSize = this->getSize();
	bSize = other.getSize();
	
	
	if(aSize[1] != bSize[0]){
		cout << "The column of the 1st matrix and the row of the 2nd matrix do not match\n";
		return Error;
	}
	
	Matris result(0, aSize[0], bSize[1]);
	int res = 0, k = 0;
	
	// i ve j yeni matrisin boyutu
	// burada k çarpýlýcak matrislerin sütun ile satýrlarýný dönüyor.
	for(int i = 0; i < aSize[0]; i++){
		for(int j = 0; j < bSize[1]; j++){
			for(int k = 0; k < bSize[0]; k++){
				result.value[i][j] += this->value[i][k] * other.value[k][j];	
			}
		}
	}
	
	return result;
}

Matris Matris::operator-(Matris &other){
	Matris Error(-1, 2, 2, true);
	int* aSize = new int[2];
	int* bSize = new int[2];
	
	aSize = this->getSize();
	bSize = other.getSize();
	
	Matris Result(0, aSize[0], aSize[1]);
	
	if(aSize[0] != bSize[0] && aSize[1] != bSize[1]){
		cout << "Size's do not match!'\n";
		return Error;
	}
	
	other.scalerMultiply(-1);
	
	for(int i = 0; i < aSize[0]; i++){
		for(int j = 0; j < aSize[1]; j++){
			Result.value[i][j] += this->value[i][j] + other.value[i][j];
		}
	}
	
	return Result;
}

Matris Matris::operator*(Double &other){
	int *size = new int[2];
	size = this->getSize();
	
	Matris result(0, size[0], size[1]);
	
	this->scalerMultiply(other.getValue());
	
	this->setValue(result);
	
	return result;
}

// Bu Fonksiyon this-> deki value'yi alýyor parametrenin value'suna koyuyor.
void Matris::setValue(Matris &other){
	int *size = new int[2];
	size = this->getSize();
	
	for(int i = 0; i < size[0]; i++){
		for(int j = 0; j < size[1]; j++){
			other.value[i][j] = this->value[i][j];
		}
	}
}

int matrisDet(arr2d value){
	int row = value.size(), col = value[0].size();
	
	
	if(row == 1 && col == 1) return value[0][0];
	if(row == 2 && col == 2){
		int left = 1, right = 1, det = 0;
		
		for(int i = 0; i < row; i++){
			for(int j = 0; j < col; j++){
				
				if(i == j){
					right *= value[i][j];
				}
				else{
					left *= value[i][j];
				}
			}
		}
		
		det = right - left;
		
		return det;
	}
	if(row > 2 && col > 2){
		vector<double> minor_row;
		arr2d minor;
		vector<arr2d> minors;
		
		int R = 0, C = 0, det = 0;
		
		while(C < col){
		
			for(int i = 0; i < row; i++){
				for(int j = 0; j < col; j++){
					if(i != R && j != C){
						minor_row.push_back(value[i][j]);
					}
				}
				if(!minor_row.empty()){
					minor.push_back(minor_row);
					minor_row.clear();
				}
			}
			
			minors.push_back(minor);
			minor.clear();
			
			C++;
		}
		
		for(int i = 0; i < minors.size(); i++){
			det += matrisDet(minors[i]) * value[0][i] * pow(-1, i);
		}
		
		return det;
		
	}
	return 0;
}

int Matris::determinant(){
	int* size = new int[2];
	size = this->getSize();
	
	if(size[0] == size[1])
		return matrisDet(this->value);
	else{
		cout << "This is not a Square Matris\n";
		return -1;
	}
}

void Matris::transpoze(){
	int* size = new int[2];
	size = this->getSize();
	
	int row = size[0], col = size[1];
	
	// Ters boyutlu açtýðýmýz array'e matrisin deðelerini koypaladýk.
	double arr[col][row];
	
	for(int i = 0; i < row; i++){
		for(int j = 0; j < col; j++){
			arr[j][i] = this->value[i][j];
		}
	}
	
	this->value.clear(); // Kendi Matrisimizi Temizledik.
	
	// Yeni Vektör Açtýk
	vector<double> newValue;
	
	// Yeni Boyutta vekötürümüze koypaladýðýmýz array'deki deðerleri atýyoruz.
	for(int i = 0; i < col; i++){
		for(int j = 0; j < row; j++){
			newValue.push_back(arr[i][j]);
		}
		this->value.push_back(newValue);
		newValue.clear();
	}
}

arr2d Matris::getValue(){
	return this->value;
}

void Matris::scalerMultiply(double value){
	int *size = new int[2];
	size = this->getSize();
	
	for(int i = 0; i < size[0]; i++){
		for(int j = 0; j < size[1]; j++){
			this->value[i][j] *= value;
		}
	}
}

void Matris::setId(int Id){
	this->id = Id;
}

int Matris::getId(){
	return this->id;
}

