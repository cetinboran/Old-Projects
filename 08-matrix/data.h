#ifndef DATA_H
#define DATA_H

#include "matris.h"


class Matris;

class Data{
	private:
		int id;
		std::vector<Matris> Matrises;
		
	public:	
		Data();
		
		void addData(Matris &other);
		void deleteData(int id);
		void listData();
		
		int* validIds();
		Matris byId(char c);
		int getSize();
		
		Matris chooseMatris(int id);
};


#endif
