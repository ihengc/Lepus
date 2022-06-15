package clause

/********************************************************
* @author: Ihc
* @date: 2022/6/15 0015 16:21
* @version: 1.0
* @description:
*********************************************************/

type Insert struct{}

func (i Insert) Name() string {
	return "INSERT"
}

func (i Insert) MergeClause() {

}
