# ask4me
generate description of source file content and directory for AI Chat


# How to use

1. Install ask4me with `go install github.com/airylinus/ask4me@latest`
2. Create a new directory at home directory with name ".ask4me"
3. Copy examples/prompts.yaml to.ask4me directory, and edit it to fit your needs.
4. Run the following command in terminal:
```
ask4me pyqt5-ui-fix
```

# Further more

You may want to use this tool by keyboard shortcut, for me at Windows 11 system I use AHK (AutoHotkey) for convenience. Here's the code:

```
^c::
If GetKeyState("j", "P")  ; 
{
    Run cmd.exe /c ask4me pyqt5-ui-fix,, Hide  ; 
}
else
{
    Send, ^c  ;
}
Return
```

Save as shutcut-ask4me.ahk and run it with ahk