package modelsWithInterface

// we are using CollectMetrics common for every metrics collection, so we declaring once and reusing everywhere

type MetricsCollection interface{
	CollectMetrics() error
}

// how  interface are initialized
// how interface are used to implement methods of different struct
// initialization of interface is missing.


//  You can’t create an instance of an interface directly, but you can make a variable of the interface type to store any value that has the needed methods.

