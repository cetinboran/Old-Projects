#include "double.h"
#include "matris.h"

Double::Double(double Value){
	this->value = Value;
}

double Double::getValue(){
	return this->value;
}

Matris Double::operator*(Matris &other){
	other.scalerMultiply(this->value);
	return other;
}


