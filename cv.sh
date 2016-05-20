
# convert potentially changed files (e.g. by gofmt from inside Eclipse) back to original format as closely as possible
# so as not to introduce too much unnecessary "changed lines".

set -x
unix2dos "$@"

# git diff "$@"

# due to symlink, has to do following instead of simply "git diff"
(cd `dirname "$@"` && git diff `basename "$@"` )

