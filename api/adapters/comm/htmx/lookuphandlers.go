package htmx

// func (ws WebServer) GetMemberDetail(c *gin.Context) {
// 	id := c.Param("id")
// 	ms := application.NewMemberService(ws.dbContext)
// 	member, err := ms.GetMember(id)
// 	if err != nil {
// 		c.HTML(500, "memberdetail.html", nil)
// 		return
// 	}
// 	c.HTML(200, "memberdetail.html", gin.H{
// 		"Member": member,
// 	})
// }