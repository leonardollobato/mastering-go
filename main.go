package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}

	arguments := os.Args
	var err error = errors.New("An error")
	k := 1
	var n float64

	for err != nil {
		if k >= len(arguments) {
			fmt.Println("None of the arguments is a float!")
			return
		}

		n, err = strconv.ParseFloat(arguments[k], 64)
		k++
	}

	min, max := n, n

	for i := 2; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			if n < min {
				min = n
			}

			if n > max {
				max = n
			}
		}

	}

	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}

//var LOGFILE = "./masteringgo.log"
/*
func error_main() {
	err := returnError(1, 2)
	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}

	err = returnError(10, 10)
	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}

	if err.Error() == "Error in returnError() function!" {
		fmt.Println("!!")
	}
}

func returnError(a, b int) error {
	if a == b {
		err := errors.New("Error in returnError() function!")
		return err
	} else {
		return nil
	}
}

func customlog_main() {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	iLog := log.New(f, "customLogLineNumber ", log.LstdFlags)

	iLog.SetFlags(log.LstdFlags | log.Lshortfile)
	iLog.Println("Hello there!")
	iLog.Println("Another log entry!")
}

func syslog_main() {
	programName := filepath.Base(os.Args[0])
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)

	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}

	log.Println("LOG_INFO + LOG_LOCAL7: Loggin in GO!")

	sysLog, err = syslog.New(syslog.LOG_MAIL, "Some program")

	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}

	log.Panic("LOG_MAIL: Logging in Go!")
	fmt.Println("Will you see this?")
}

func stdout_main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}

	arguments := os.Args
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)

	for i := 2; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}

	}

	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}

func stdin_main() {
	f := os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}

func sdout_main() {
	my_string := ""
	arguments := os.Args

	if len(arguments) == 1 {
		my_string = "Please give me one argument!"
	} else {
		my_string = arguments[1]
	}

	io.WriteString(os.Stdout, my_string)
	io.WriteString(os.Stdout, "\n")
}*/
