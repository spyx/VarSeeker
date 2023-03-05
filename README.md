# VarSeeker 

VarSeeker is small utility which get list of url from stdin and search from all possible variables inside. This is good for search hidden parameters. For example some javascript file can use some variable from GET parameter and this parameter is being reflected back. This scrip only search for it. 

/home/kali/Projects/ParamAngler/README.md/home/kali/Projects/ParamAngler/README.md


# Instalation

With go you can use this command

```
go install github.com/spyx/ParamVarSeekerAngler@latest
```

# Usage

```
go run main.go -h                     


$$\    $$\                    $$$$$$\                      $$\                           
$$ |   $$ |                  $$  __$$\                     $$ |                          
$$ |   $$ |$$$$$$\   $$$$$$\ $$ /  \__| $$$$$$\   $$$$$$\  $$ |  $$\  $$$$$$\   $$$$$$\  
\$$\  $$  |\____$$\ $$  __$$\\$$$$$$\  $$  __$$\ $$  __$$\ $$ | $$  |$$  __$$\ $$  __$$\ 
 \$$\$$  / $$$$$$$ |$$ |  \__|\____$$\ $$$$$$$$ |$$$$$$$$ |$$$$$$  / $$$$$$$$ |$$ |  \__|
  \$$$  / $$  __$$ |$$ |     $$\   $$ |$$   ____|$$   ____|$$  _$$<  $$   ____|$$ |      
   \$  /  \$$$$$$$ |$$ |     \$$$$$$  |\$$$$$$$\ \$$$$$$$\ $$ | \$$\ \$$$$$$$\ $$ |      
        \_/    \_______|\__|      \______/  \_______| \_______|\__|  \__| \_______|\__|                               


Remember that bug bounty and security tools should only be used ethically and responsibly.
Misuse of these tools can lead to harm and legal consequences.
Use these tools with caution and obtain permission before performing any testing or analysis.

Usage: VarSeeker [OPTIONS]
Options:
  -f string
        path to a file containing a list of URLs
  -h    display usage information
  -s    hide banner
  -t int
        Number of goroutines to use (default 1)

```


