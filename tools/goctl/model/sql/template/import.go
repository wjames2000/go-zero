package template

var (
	Imports = `import (
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"github.com/wjames2000/go-zero/core/stores/cache"
	"github.com/wjames2000/go-zero/core/stores/sqlc"
	"github.com/wjames2000/go-zero/core/stores/sqlx"
	"github.com/wjames2000/go-zero/core/stringx"
	"github.com/wjames2000/go-zero/tools/goctl/model/sql/builderx"
)
`
	ImportsNoCache = `import (
	"database/sql"
	"strings"
	{{if .time}}"time"{{end}}

	"github.com/wjames2000/go-zero/core/stores/sqlc"
	"github.com/wjames2000/go-zero/core/stores/sqlx"
	"github.com/wjames2000/go-zero/core/stringx"
	"github.com/wjames2000/go-zero/tools/goctl/model/sql/builderx"
)
`
)
