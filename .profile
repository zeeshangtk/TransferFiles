export HADOOP_INSTALL=/Users/zeeshana/Code/lib/hadoop/hadoop-2.7.1/
export PATH="/Users/zeeshana/Code/lib/scala-2.10.4/bin:$HADOOP_INSTALL/bin/:$PATH"

export PATH="$PATH:/Users/zeeshana/Code/play2/activator-1.3.6-minimal/"
export PATH="$PATH:/usr/local/Cellar/smlnj/110.78/bin/:/Users/zeeshana/go-lang/bin"
export CLASSPATH=".:/usr/local/lib/antlr-4.5-complete.jar:$CLASSPATH"


alias gpr='git pull --rebase'
alias ga='git add'
alias gp='git push'

export GOPATH="/Users/zeeshana/Code/go-lang/"
export GO15VENDOREXPERIMENT=1



#Git specific
alias gs="git status"
alias gd="git diff"
alias gco="git checkout"
alias gcm="git commit"
alias gpr="git pull -r"
alias gpsh="git push"
alias ga="git add"
alias gl="git log"
alias glo="git log --oneline --decorate --all --graph"
alias gst="git stash"
alias gr="git rebase"
alias gcp="git cherry-pick"
alias grh="git reset --hard"
alias gst='git status'
alias grs="git reset --soft"
alias gsu="git submodule update"
alias gg="git gui &"
alias gfb="gulp format-js-before-checkin build"
alias gb="gulp build"
