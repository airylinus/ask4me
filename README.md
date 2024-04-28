# ask4me
generate description of source file content and directory for AI Chat


# How to use

1. Install ask4me with `go install github.com/airylinus/ask4me@latest`
2. Create a new directory at home directory with name ".ask4me"
3. Copy examples/prompts.yaml to.ask4me directory
4. Run the following command in terminal:
```
ask4me pyqt4-ui-fix
```

# Further more

You may want using this tool by shortcut, for me at Windows 11 system i use AHK (AutoHotkey) to create a hotkey to run the command. Here's the code:

```
^c::
If GetKeyState("j", "P")  ; 检查是否按下 J 键
{
    Run cmd.exe /c ask4me pyqt5-ui-fix,, Hide  ; 
}
else
{
    Send, ^c  ;
}
Return

save as shutcut-ask4me.ahk and run it with ahk