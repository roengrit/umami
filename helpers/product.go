package helpers

import (
	"bytes"
	"strconv"
	"strings"
	m "umami/models"
)

//HTMLProTemplate _
const HTMLProTemplate = `<tr>
							<td>{name}</td>
							<td>{qty}</td>
							<td>{category}</td> 
							<td>{status}</td> 
							<td>
								<div class="btn-group">
									{action}
								</div>
							</td>                             
						</tr>`

//HTMLProActionEnable _
const HTMLProActionEnable = `<a class="btn bg-purple" title="รายละเอียด" target="_blank" href="/product/movement/?id={id}"><i class="fa fa-file-text-o"></i></a>
							 <a class="btn btn-primary " title="แก้ไข"  href="/product/?id={id}"><i class="fa fa-edit"></i></a>
							 <button type="button" class="btn btn-primary dropdown-toggle" data-toggle="dropdown" aria-expanded="false">
										<span class="caret"></span>
										<span class="sr-only">Toggle Dropdown</span>
									</button>
									<ul class="dropdown-menu" role="menu">
									    <li><a target="_blank" href="/product/formular/?id={id}" title="วัตถุดิบ">วัตถุดิบ</a></li>
										<li><a href="#" onclick='confirmDeleteGlobal({id},"/product")' title="ลบ">ลบ</a></li>
								</ul> `

//HTMLProNotFoundRows _
const HTMLProNotFoundRows = `<tr> <td  colspan="4" style="text-align:center;">*** ไม่พบข้อมูล ***</td></tr>`

//HTMLProError _
const HTMLProError = `<tr> <td  colspan="4" style="text-align:center;">{err}</td></tr>`

//GenProHTML _
func GenProHTML(lists []m.Product) string {
	var hmtlBuffer bytes.Buffer
	for _, val := range lists {
		temp := strings.Replace(HTMLProTemplate, "{name}", val.Name, -1)
		temp = strings.Replace(temp, "{qty}", RenderFloat("#,###.##", val.BalanceQty), -1)
		temp = strings.Replace(temp, "{category}", val.ProductCategory.Name, -1)
		tempAction := strings.Replace(HTMLProActionEnable, "{id}", strconv.Itoa(val.ID), -1)
		if val.Active {
			temp = strings.Replace(temp, "{status}", "เปิดใช้งาน", -1)
		} else {
			temp = strings.Replace(temp, "{status}", "ปิดใช้งาน", -1)
		}
		temp = strings.Replace(temp, "{action}", tempAction, -1)
		hmtlBuffer.WriteString(temp)
	}
	return hmtlBuffer.String()
}

//HTMLProCateTemplate _
const HTMLProCateTemplate = `<tr>
							<td>{name}</td>
							<td>
								<div class="btn-group">
									{action}
								</div>
							</td>                             
						</tr>`

//HTMLProCateActionEnable _
const HTMLProCateActionEnable = ` <a class="btn btn-primary"   title="แก้ไข"  href="/product-category/?id={id}"><i class="fa fa-edit"></i></a>
								  <a class="btn btn-danger"   title="ลบ" href="#"  onclick='confirmDeleteGlobal({id},"/product-category")'><i class="fa fa-trash-o"></i></a>`

//HTMLProCateNotFoundRows _
const HTMLProCateNotFoundRows = `<tr> <td  colspan="2" style="text-align:center;">*** ไม่พบข้อมูล ***</td></tr>`

//HTMLProCateError _
const HTMLProCateError = `<tr> <td  colspan="2" style="text-align:center;">{err}</td></tr>`

//GenProCateHTML _
func GenProCateHTML(lists []m.ProductCategory) string {
	var hmtlBuffer bytes.Buffer
	for _, val := range lists {
		temp := strings.Replace(HTMLProCateTemplate, "{name}", val.Name, -1)
		tempAction := strings.Replace(HTMLProCateActionEnable, "{id}", strconv.Itoa(val.ID), -1)
		temp = strings.Replace(temp, "{action}", tempAction, -1)
		hmtlBuffer.WriteString(temp)
	}
	return hmtlBuffer.String()
}
