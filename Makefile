build:
	/usr/bin/env bash make.sh
back:
	/usr/bin/env bash make.sh back $(version)
version:
	/usr/bin/env bash make.sh version
list:
	/usr/bin/env bash make.sh list
publish:
	/usr/bin/env bash make.sh publish
test:
	/usr/bin/env bash make.sh test
upgrade:
	/usr/bin/env bash make.sh upgrade