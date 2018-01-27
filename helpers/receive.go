package helpers

import (
	"bytes"
	"strconv"
	"strings"
	m "umami/models"
)

//HTMLReceiveTemplate _
const HTMLReceiveTemplate = `<tr>
							<td>{doc_date}</td>
							<td>{doc_no}</td>
							<td>{supplier_name}</td>
							<td>{total_net_amount}</td> 
							<td>
								<div class="btn-group">
									{action}
								</div>
							</td>                             
						</tr>`

//HTMLReceiveActionEnable _
const HTMLReceiveActionEnable = `<a class="btn bg-purple" title="รายละเอียด" target="_blank" href="/receive/read/?id={id}"><i class="fa fa-file-text-o"></i></a>
							     <a class="btn btn-primary " title="แก้ไข"  href="/receive/?id={id}"><i class="fa fa-edit"></i></a>`

//HTMLReceiveNotFoundRows _
const HTMLReceiveNotFoundRows = `<tr> <td  colspan="5" style="text-align:center;">*** ไม่พบข้อมูล ***</td></tr>`

//HTMLReceiveError _
const HTMLReceiveError = `<tr> <td  colspan="5" style="text-align:center;">{err}</td></tr>`

//GenReceiveHTML _
func GenReceiveHTML(lists []m.Receive) string {
	var hmtlBuffer bytes.Buffer
	for _, val := range lists {
		temp := strings.Replace(HTMLReceiveTemplate, "{doc_date}", val.DocDate.Format("02-01-2006"), -1)
		temp = strings.Replace(temp, "{doc_no}", val.DocNo, -1)
		temp = strings.Replace(temp, "{supplier_name}", val.SupplierName, -1)
		temp = strings.Replace(temp, "{total_net_amount}", ThCommaSep(val.TotalNetAmount), -1)
		tempAction := strings.Replace(HTMLReceiveActionEnable, "{id}", strconv.Itoa(val.ID), -1)
		temp = strings.Replace(temp, "{action}", tempAction, -1)
		hmtlBuffer.WriteString(temp)
	}
	return hmtlBuffer.String()
}
