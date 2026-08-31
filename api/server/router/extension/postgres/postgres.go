package postgres

import (
	"github.com/caoyingjunz/pixiu/api/server/httputils"
	"github.com/caoyingjunz/pixiu/api/server/router/apiregistry"
	"github.com/caoyingjunz/pixiu/cmd/app/options"
	"github.com/caoyingjunz/pixiu/pkg/controller"
	"github.com/caoyingjunz/pixiu/pkg/types"
	"github.com/gin-gonic/gin"
)

type router struct{ c controller.PixiuInterface }

type meta struct {
	DatasourceId int64 `uri:"datasourceId"`
}

type opts struct {
	Database string `form:"database"`
	Schema   string `form:"schema"`
}

type sessionOpts struct {
	PID       int64 `form:"pid" binding:"required"`
	Terminate bool  `form:"terminate"`
}

func RegisterPostgres(o *options.Options, g *apiregistry.Group) {
	r := &router{c: o.Controller}
	g.Entries = append(g.Entries,
		apiregistry.RouteEntry{Method: "POST", RelativePath: "/postgres/ping", Handler: r.pingAdhoc, Description: "PostgreSQL 临时连通性探测"},
		apiregistry.RouteEntry{Method: "GET", RelativePath: "/postgres/:datasourceId/ping", Handler: r.ping, Description: "PostgreSQL 连接探测"},
		apiregistry.RouteEntry{Method: "GET", RelativePath: "/postgres/:datasourceId/info", Handler: r.info, Description: "PostgreSQL 实例概览"},
		apiregistry.RouteEntry{Method: "GET", RelativePath: "/postgres/:datasourceId/databases", Handler: r.databases, Description: "PostgreSQL 数据库列表"},
		apiregistry.RouteEntry{Method: "GET", RelativePath: "/postgres/:datasourceId/schemas", Handler: r.schemas, Description: "PostgreSQL Schema 列表"},
		apiregistry.RouteEntry{Method: "GET", RelativePath: "/postgres/:datasourceId/tables", Handler: r.tables, Description: "PostgreSQL 表列表"},
		apiregistry.RouteEntry{Method: "POST", RelativePath: "/postgres/:datasourceId/query", Handler: r.query, Description: "PostgreSQL SQL 控制台"},
		apiregistry.RouteEntry{Method: "GET", RelativePath: "/postgres/:datasourceId/sessions", Handler: r.sessions, Description: "PostgreSQL 会话列表"},
		apiregistry.RouteEntry{Method: "DELETE", RelativePath: "/postgres/:datasourceId/sessions", Handler: r.cancel, Description: "PostgreSQL 取消或终止会话"})
}

func (r *router) pingAdhoc(c *gin.Context) {
	var q types.PostgresSourceConfig
	res := httputils.NewResponse()
	if e := c.ShouldBindJSON(&q); e != nil {
		httputils.SetFailed(c, res, e)
		return
	}
	result, e := r.c.Extension().Postgres().PingAdhoc(c, &q)
	if e != nil {
		httputils.SetFailed(c, res, e)
		return
	}
	res.Result = result
	httputils.SetSuccess(c, res)
}

func (r *router) ping(c *gin.Context) {
	var m meta
	res := httputils.NewResponse()
	if e := httputils.ShouldBindAny(c, nil, &m, nil); e != nil {
		httputils.SetFailed(c, res, e)
		return
	}
	result, e := r.c.Extension().Postgres().Ping(c, m.DatasourceId)
	if e != nil {
		httputils.SetFailed(c, res, e)
		return
	}
	res.Result = result
	httputils.SetSuccess(c, res)
}

func (r *router) info(c *gin.Context) {
	var m meta
	res := httputils.NewResponse()
	if e := httputils.ShouldBindAny(c, nil, &m, nil); e != nil {
		httputils.SetFailed(c, res, e)
		return
	}
	result, e := r.c.Extension().Postgres().Info(c, m.DatasourceId)
	if e != nil {
		httputils.SetFailed(c, res, e)
		return
	}
	res.Result = result
	httputils.SetSuccess(c, res)
}

func (r *router) databases(c *gin.Context) {
	var m meta
	res := httputils.NewResponse()
	if e := httputils.ShouldBindAny(c, nil, &m, nil); e != nil {
		httputils.SetFailed(c, res, e)
		return
	}
	result, e := r.c.Extension().Postgres().ListDatabases(c, m.DatasourceId)
	if e != nil {
		httputils.SetFailed(c, res, e)
		return
	}
	res.Result = result
	httputils.SetSuccess(c, res)
}

func (r *router) schemas(c *gin.Context) {
	var m meta
	var o opts
	res := httputils.NewResponse()
	if e := httputils.ShouldBindAny(c, nil, &m, &o); e != nil {
		httputils.SetFailed(c, res, e)
		return
	}
	result, e := r.c.Extension().Postgres().ListSchemas(c, m.DatasourceId, o.Database)
	if e != nil {
		httputils.SetFailed(c, res, e)
		return
	}
	res.Result = result
	httputils.SetSuccess(c, res)
}

func (r *router) tables(c *gin.Context) {
	var m meta
	var o opts
	res := httputils.NewResponse()
	if e := httputils.ShouldBindAny(c, nil, &m, &o); e != nil {
		httputils.SetFailed(c, res, e)
		return
	}
	result, e := r.c.Extension().Postgres().ListTables(c, m.DatasourceId, o.Database, o.Schema)
	if e != nil {
		httputils.SetFailed(c, res, e)
		return
	}
	res.Result = result
	httputils.SetSuccess(c, res)
}

func (r *router) query(c *gin.Context) {
	var m meta
	var q types.PostgresQueryRequest
	res := httputils.NewResponse()
	if e := httputils.ShouldBindAny(c, &q, &m, nil); e != nil {
		httputils.SetFailed(c, res, e)
		return
	}
	result, e := r.c.Extension().Postgres().ExecuteSQL(c, m.DatasourceId, &q)
	if e != nil {
		httputils.SetFailed(c, res, e)
		return
	}
	res.Result = result
	httputils.SetSuccess(c, res)
}

func (r *router) sessions(c *gin.Context) {
	var m meta
	res := httputils.NewResponse()
	if e := httputils.ShouldBindAny(c, nil, &m, nil); e != nil {
		httputils.SetFailed(c, res, e)
		return
	}
	result, e := r.c.Extension().Postgres().ListSessions(c, m.DatasourceId)
	if e != nil {
		httputils.SetFailed(c, res, e)
		return
	}
	res.Result = result
	httputils.SetSuccess(c, res)
}

func (r *router) cancel(c *gin.Context) {
	var m meta
	var o sessionOpts
	res := httputils.NewResponse()
	if e := httputils.ShouldBindAny(c, nil, &m, &o); e != nil {
		httputils.SetFailed(c, res, e)
		return
	}
	if e := r.c.Extension().Postgres().CancelSession(c, m.DatasourceId, o.PID, o.Terminate); e != nil {
		httputils.SetFailed(c, res, e)
		return
	}
	httputils.SetSuccess(c, res)
}
