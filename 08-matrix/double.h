#ifndef DOUBLE_H
#define DOUBLE_H

#include "matris.h"

class Matris;

class Double{
	private:
		double value;
	
	public:
		Double(double Value = 0);
		
		double getValue();
		Matris operator*(Matris &other);
};


#endif
