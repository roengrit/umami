package helpers

import (
	"bytes"
	"strconv"
	"strings"
	m "umami/models"
)

//HTMLStockPickCountTemplate _
const HTMLStockPickCountTemplate = `<tr>
							<td>{doc_date}</td>
							<td>{doc_no}</td>	
							<td>{remark}</td>					 
							<td>{active}</td>
							<td>{total_net_amount}</td> 
							<td>
								<div class="btn-group">
									{action}
								</div>
							</td>                             
						</tr>`

//HTMLStockCountActionEnable _
const HTMLStockCountPickActionEnable = `<a class="btn bg-purple" title="รายละเอียด" target="_blank" href="/stock/read/?id={id}"><i class="fa fa-file-text-o"></i></a>
								 <a class="btn btn-primary " title="แก้ไข"  href="/stock/?id={id}"><i class="fa fa-edit"></i></a>
								 <button type="button" class="btn btn-primary dropdown-toggle" data-toggle="dropdown" aria-expanded="false">
										<span class="caret"></span>
										<span class="sr-only">Toggle Dropdown</span>
									</button>
									<ul class="dropdown-menu" role="menu">
										<li><a href="#" onclick="cancelDoc({id})" title="ยกเลิก">ยกเลิก</a></li>
								</ul> `

//HTMLStockCountActionEditOnly _
const HTMLStockCountPickActionEditOnly = `<a class="btn bg-purple" title="รายละเอียด" target="_blank" href="/stock/read/?id={id}"><i class="fa fa-file-text-o"></i></a>
								   <a class="btn btn-primary" title="แก้ไข"  target="_blank" href="/stock/?id={id}"><i class="fa fa-edit"></i></a>
								   <button type="button" class="btn btn-primary dropdown-toggle" data-toggle="dropdown" aria-expanded="false">
										<span class="caret"></span>
										<span class="sr-only">Toggle Dropdown</span>
								   </button>
								  `

//HTMLStockCountNotFoundRows _
const HTMLStockCountPickNotFoundRows = `<tr> <td  colspan="4" style="text-align:center;">*** ไม่พบข้อมูล ***</td></tr>`

//HTMLStockCountError _
const HTMLStockCountPickError = `<tr> <td  colspan="4" style="text-align:center;">{err}</td></tr>`

//GenStockCountHTML _
func GenStockCountPickHTML(lists []m.StockCountPick) string {
	var hmtlBuffer bytes.Buffer
	for _, val := range lists {
		temp := strings.Replace(HTMLStockPickCountTemplate, "{doc_date}", val.DocDate.Format("02-01-2006"), -1)
		temp = strings.Replace(temp, "{doc_no}", val.DocNo, -1)
		temp = strings.Replace(temp, "{remark}", val.Remark, -1)
		if val.FlagTemp == 1 {
			temp = strings.Replace(temp, "{active}", "W", -1)
		} else {
			if val.Active {
				temp = strings.Replace(temp, "{active}", "N", -1)
			} else {
				temp = strings.Replace(temp, "{active}", "C", -1)
			}
		}
		temp = strings.Replace(temp, "{total_net_amount}", ThCommaSep(val.TotalNetAmount), -1)
		if val.Active {
			tempAction := strings.Replace(HTMLStockCountPickActionEnable, "{id}", strconv.Itoa(val.ID), -1)
			temp = strings.Replace(temp, "{action}", tempAction, -1)
		} else {
			tempAction := strings.Replace(HTMLStockCountPickActionEditOnly, "{id}", strconv.Itoa(val.ID), -1)
			temp = strings.Replace(temp, "{action}", tempAction, -1)
		}
		hmtlBuffer.WriteString(temp)
	}
	return hmtlBuffer.String()
}
