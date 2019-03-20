

## 关于本文·16.10.2小结糟糕错误处理的一些见解

本文仅表达译者对错误处理的观点，并且觉得原文说的并不很合理，希望不会误导（我个人观点）其他入门读者。

### 关于16.10.2的第一个代码示例

16.10.2小结中关于错误处理的第一个代码示例是标准且通用的错误处理方式。
文中认为这种错误处理方式会使你的代码中充满`if err != nil {...}`，认为这样会令人难以分辨正常的程序逻辑与错误处理（难道错误处理不算做正常的程序逻辑么:)）。

**书中代码示例一**：

```Go
... err1 := api.Func1()
if err1 != nil {
    fmt.Println("err: " + err.Error())
    return
}
err2 := api.Func2()
if err2 != nil {
...
    return
}
```

**我的观点**：

1、错误处理也是正常程序逻辑的一部分，程序逻辑不就是对一个操作可能出现的结果进行判断，
并对每一种结果做相应的后续处理么。错误是我们已知的可能会出现的一种结果，我们也需要处理这种情况，它也是正常逻辑的一部分。显然，把错误单独拎出来，与正常逻辑并列来做对待，并不合理。


2、在其他语言中，我们可能会用到 try... catch...语句来对可能出现的错误进行处理，难道你会说try-catch语句让你的代码一团糟，程序逻辑和错误处理混在一起很复杂，让你阅读代码困难么。绝大多数情况下，让你感觉难以阅读甚至恶心（可能形容过度了）的代码绝不会是因为错误处理相关的代码导致的，而是当时写这些代码的人逻辑不清甚至逻辑混乱造成的。


3、这个可能和每个人的习惯（自己写代码的思路、风格）或者说适应（看其他人的代码时能很快习惯作者的代码风格）有关，我每次看代码都会先略过错误处理的部分，那么剩下的就是理想情况下的程序逻辑了，如果对某一处心存疑惑那么就再仔细看这部分的代码。毕竟我们写的代码绝大多数情况下是希望它按理想的情况跑的，

_ _ _

### 关于16.10.2的第二个代码示例

16.10.2小结中关于错误处理的第二个代码示例是推荐给我们的错误处理方式，对于其推荐的这种方式，个人认为是有一定的适用范围的，并不适合大多数的错误处理，反而在处理某些业务逻辑时可以使用，比如将不符合业务逻辑的情况视作一种错误（自定义）来统一做处理。

**书中代码示例二**：

```Go
func httpRequestHandler(w http.ResponseWriter, req *http.Request) {
    err := func () error {
        if req.Method != "GET" {
            return errors.New("expected GET")
        }
        if input := parseInput(req); input != "command" {
            return errors.New("malformed command")
        }
        // 可以在此进行其他的错误检测
    } ()

        if err != nil {
            w.WriteHeader(400)
            io.WriteString(w, err)
            return
        }
        doSomething() ...
```

1、代码示例二中对不符合业务逻辑的两种情况做了归类，并自定义了错误，做了统一的处理。这样从业务层面来看，将不符合业务逻辑的情况视为错误，统一写到了匿名函数中，剩下了一个统一的错误处理与正常的业务逻辑。或许采用这种方式处理这类场景还不错，但是如果换作下面的这个示例可能就不是很合理了。

下面的示例一是采用了作者推荐的统一处理错误方式，示例二使用的是通常的错误处理方式

**示例一**：

```Go
// 目标目录下包含多种Archive格式文件，将其中的'x-msdownload'类型文件移动到其他目录下
func moveEXE(files []os.FileInfo, aimPath, exePath string) {
	var numExe, numOther int
	var fileBuf []byte
	var fileType types.Type

	for _, file := range files {
		fileName := aimPath + file.Name()
		newFileName := exePath + file.Name()

		err := func() error {  
            
            // 读取文件内容
			if buf, err := ioutil.ReadFile(fileName); err != nil {
				log.Printf("Time of read file: %s occur error: %s\n", fileName, err)
				return err
			}else {
				fileBuf = buf
			}
            
            // 判断文件是否为Archive（压缩）格式
			if kind, err := filetype.Archive(fileBuf); err!= nil {
				log.Printf("Time of judge file type occur error: %s\n", err)
				return err
			}else {
				fileType = kind
			}
            
            // 文件是否为'x-msdownload'类型
			if fileSubType := fileType.MIME.Subtype; fileSubType == "x-msdownload" {
				log.Printf("file : %s is exe file\n", fileName)
				if err := os.Rename(fileName, newFileName); err != nil {
					log.Printf("mv file: %s faile, error is: %s\n", fileName, err)
					return err
				}
				numExe ++
			}else {
				log.Println("no exe")
				numOther ++
			}
			return nil
		}()

		if err != nil {
			continue
		}
    }
    log.Printf("exe file num is: %d, other file num is: %d", numExe, numOther)
}
```

1、通常来说，我们使用匿名函数是因为部分操作不值得新定义一个函数或者该函数仅使用一次，示例一中的匿名函数包含了很多操作，或许我们应该为此重新定义一个函数。其中包含了几乎全部的逻辑代码，我想这看起来并不是啥好主意，甚至如果你把更多的逻辑代码放到了匿名函数里，看起来应该会更加糟糕。

**示例二**：

```Go

// 目标目录下包含多种Archive格式文件，将其中的'x-msdownload'类型文件移动到其他目录下
func moveEXE(files []os.FileInfo, aimPath, exePath string) {
	var numExe, numOther int

	for _, file := range files {
		fileName := aimPath + file.Name()
		newFileName := exePath + file.Name()
		
        // 读取文件内容
		buf, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Printf("read file:%s  occur error\n", fileName)
			continue
		}
		
        // 判断文件是否为Archive（压缩）格式
		kind, err := filetype.Archive(buf)
		if err != nil {
			log.Println("judge file type error")
			continue
		}

        // 获取文件具体的类型
		fileSubType := kind.MIME.Subtype

        // 文件是否为'x-msdownload'类型
		if fileSubType == "x-msdownload" {
			log.Printf("file : %s is exe file\n", fileName)
			err := os.Rename(fileName, newFileName)
			if err != nil {
				log.Printf("mv file: %s faile\n", fileName)
				continue
			}
			numExe ++
		}else {
			log.Println("no exe")
			numOther ++
		}
	}
	log.Printf("exe file num is: %d, other file num is: %d", numExe, numOther)
}
```

2、示例二中的代码看起来则自然多了（我是这种感觉），或许你认为这俩个例子相差无几，但是我想通过他们表明，原文16.10.2中推荐的错误处理方式是有一定的使用场景的，并不能取代标准且通用的错误处理方式，希望大家能够注意。

---



### 关于错误处理的一些延伸

1、除了使用Go中已经定义好的error，我们也可以根据需要自定义error。

下面的示例三，我们自定义了parseError 错误，展示了发生错误的文件和具体的错误信息，在你读取目录下的多个文件时可以方便的告诉你具体在读哪个文件时发生了错误（作为示例，仅读取单个文件）。

示例四中，展示了调用 parseFile 函数时，调用者可以采用的一种错误处理方式，根据错误的类型，采取对应的操作。

**示例三**：

```go

type parseError struct {
	File *os.File
	ErrorInfo string
}


func (e *parseError) Error() string {
	errInfo := fmt.Sprintf(
		"parse file: %s occur error, error info: %s",
		e.File.Name(),
		e.ErrorInfo)
	return errInfo
}


func parseFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	var buf [512]byte
	for {
		switch num, err := f.Read(buf[:]); {
		case num < 0:
			readError := parseError{f, err.Error()}
             log.Println(readError.Error())
			return &readError
            
		case num == 0:
			readError := parseError{f, err.Error()}
             log.Println(readError.Error())
			return &readError

		case num > 0:
			fmt.Println(string(buf[:num]))
             log.Printf("read file: %s contents normally")
		}
	}
}

```

**示例四**：

```go
func main()  {
	err := parseFile("/home/rabbit/go/test_use/test")
	switch err := err.(type) {
        
	case *parseError:
        log.Println("parse error: ", err)
        
	case *os.PathError:
        log.Println("path error: ", err)
	}
}
```

2、如果你想在返回错误之前做一些额外的操作，比如记录日志，那你可以单独写一个额外处理错误的函数或者一个匿名函数就可以（这取决于你是否常用该函数或它的功能是否很多），类似Python中的装饰器一样。

示例五中，handleError 将错误写入到了指定日志文件中；

示例六中，parseFile 中使用 `defer func() {handleError("/home/rabbit/go/test_use/log", err)}()`代替了多次出现的`log.Println(readError.Error())`，并将日志记录持久化到文件中。

**示例五**:

```go
func handleError(logPath string, err error) {
    if err == nil {
        return
    }
    
    logFile, _ := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 666)
	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetPrefix("[FileError]")
	log.SetFlags(log.Llongfile|log.Ldate|log.Ltime)
	log.Println(err.Error())
}
```

**示例六**:

```go
func parseFile(path string) (err error) {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	defer func() {handleError("/home/rabbit/go/test_use/log", err)}()

	var buf [512]byte
	for {
		switch num, err := f.Read(buf[:]); {

		case num < 0:
			err := &parseError{f, err.Error()}
			return err

		case num == 0:
			err := &parseError{f, err.Error()}
			return err

		case num > 0:
			fmt.Println(string(buf[:num]))
		}
	}
}
```

