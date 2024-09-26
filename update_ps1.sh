update_ps1() {
    current_dir=$(pwd)
    result=$(/path/to/ps1.exe "$current_dir")
    PS1="$result "
}

PROMPT_COMMAND=update_ps1
