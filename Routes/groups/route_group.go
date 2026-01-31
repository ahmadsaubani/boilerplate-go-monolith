package groups

type apiGroup struct {
	ArticleGroup iArticleGroup
}

var RoutesGroupCollection = apiGroup{
	ArticleGroup: newArticleRouterGroup(),
}
