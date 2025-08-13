#include "data.h"
#include "matris.h"


Data::Data(){
	this->id = 1;
}

void Data::addData(Matris &other){
	other.setId(this->id);
	this->id++;
	this->Matrises.push_back(other);
}

void Data::deleteData(int id){
	int size = this->getSize();
	
	for(int i = 0; i < size; i++){
		if(this->Matrises[i].getId() == id){
			this->Matrises.erase(this->Matrises.begin() + 1);
		}
	}
}

void Data::listData(){
	int size = this->getSize();
	
	for(int i = 0; i < size; i++){
		Matrises[i].listMatris();
	}
}

int* Data::validIds(){
	int size = this->getSize();
	int* ids = new int[size];
	
	for(int i = 0; i < size; i++){
		ids[i] = this->Matrises[i].getId();
	}
	
	return ids;
}

Matris Data::chooseMatris(int id){
	int size = this->getSize();
	Matris Error(-1, 1, 1, true);
	
	for(int i = 0; i < size; i++){
		if(this->Matrises[i].getId() == id){
			return this->Matrises[i];
		}
	}
	
	return Error;
}

int Data::getSize(){
	return this->Matrises.size();
}

Matris Data::byId(char c){
	// 97 -> a 122 -> z
	
	Matris Error(-1, 1, 1, true);
	
	int size = this->getSize(); // Data'daki matris sayisi
	int *validIds = new int[size]; // valid id'leri tutucak arr
	
	validIds = this->validIds();
	
	for(int i = 0; i < size; i++){
		if(((int) c - 96) == validIds[i]){
			return this->chooseMatris(validIds[i]);
		}
	}
	
	return Error;
}
