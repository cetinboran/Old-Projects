#ifndef FREE_H
#define FREE_H

#include "matris.h"
#include "data.h"

class Free{
	private:
		Data data;
	
	public:
		Free(Data &data);
	
		void loop();
		
		// Commands
		void help();
};

#endif
