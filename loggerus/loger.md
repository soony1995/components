Logrus has seven logging levels: Trace, Debug, Info, Warning, Error, Fatal and Panic.

- log.Trace("Something very low level.")
- log.Debug("Useful debugging information.")
- log.Info("Something noteworthy happened!")
- log.Warn("You should probably take a look at this.")
- log.Error("Something failed but I'm not quitting.")
- // Calls os.Exit(1) after logging
- log.Fatal("Bye.")
- // Calls panic() after logging
- log.Panic("I'm bailing.")

<hr>

## io.Writer vs os.Stdout

### io.Writer를 사용하는 경우:  
>io.Writer는 출력을 기록하는 인터페이스로, 파일, 버퍼, 네트워크 연결 등 다양한 출력 대상에 대해 추상화된 인터페이스를 제공합니다.  
io.Writer를 구현한 객체를 사용하여 로그를 원하는 위치에 기록할 수 있습니다. 이는 유연성과 확장성을 제공하며, 로그 메시지를 파일뿐만 아니라 다른 대상에도 전송할 수 있도록 합니다.   

### os.Stdout을 사용하는 경우: 
> os.Stdout은 Go 언어에서 표준 출력(stdout)을 나타내는 io.Writer의 구현체입니다.  
os.Stdout은 콘솔 또는 터미널 창에 출력을 보여주는 표준 출력 대상으로 사용됩니다.  
로그 메시지를 os.Stdout에 전달하면, 해당 메시지는 터미널에 출력됩니다.  

즉, **io.Writer를 사용하면 로그 메시지를 파일, 네트워크 연결 등 다양한 출력 대상으로 전송할 수 있으며, os.Stdout을 사용하면 로그 메시지가 터미널에 표시됩니다.** 선택은 개발자가 로그를 어디에 출력하고 싶은지에 따라 다를 수 있습니다.
