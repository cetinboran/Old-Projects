#ifndef MATRIS_H
#define MATRIS_H

#include <iostream>
#include <vector>
#include "double.h"
#include "data.h"

using namespace std;

typedef vector<vector<double>> arr2d;

class Double;
class Data;

class Matris{
	private:
		int id;
		arr2d value;
		
	public:
		Matris(int size);
		Matris(arr2d Value);
		Matris(double number, int row, int col, bool error = false);
		
		Matris operator+(Matris &other);
		Matris operator*(Matris &other);
		Matris operator*(Double &other);
		Matris operator-(Matris &other);
	
		static void takeFromList(Data &data);
		void printMatris();
		void listMatris();
		void transpoze();
		int* getSize();
		int determinant();
		
		arr2d getValue();
		void setValue(Matris &other); // kendi içindeki value'yu other'a veriyor.
		void scalerMultiply(double value); 
		
		void setId(int Id);
		int getId();
		
	
};


#endif
