#!/bin/sh

go install github.com/ipfn/ipfn/tools/ipfn-precommit

echo "#!/bin/sh" > .git/hooks/pre-commit
echo "ipfn-precommit" >> .git/hooks/pre-commit
