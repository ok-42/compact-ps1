update_ps1() {
    current_dir=$(pwd)
    result=$(/path/to/ps1.exe "$current_dir")
    size=${#result}
    if [ "$size" -le 20 ]; then
        PS1='\012'$ORIG_PS1
    else
        PS1="$result "
    fi
}

PROMPT_COMMAND=update_ps1
