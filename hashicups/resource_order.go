package hashicups

import (
	"context"
	"strconv"

	hc "github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOrder() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOrderCreate,
		ReadContext:   resourceOrderRead,
		UpdateContext: resourceOrderUpdate,
		DeleteContext: resourceOrderDelete,
		Schema: map[string]*schema.Schema{
			"items": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"coffee": {
							Type:     schema.TypeList,
							MaxItems: 1, // Nested Object
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"teaser": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"price": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"image": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"quantity": {
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceOrderCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*hc.Client)

	var diags diag.Diagnostics

	items := d.Get("items").([]interface{})
	ois := []hc.OrderItem{}

	for _, item := range items {
		i := item.(map[string]interface{})

		co := i["coffee"].([]interface{})[0]
		coffee := co.(map[string]interface{})

		oi := hc.OrderItem{
			Coffee: hc.Coffee{
				ID: coffee["id"].(int),
			},
			Quantity: i["quantity"].(int),
		}
		ois = append(ois, oi)
	}

	o, err := c.CreateOrder(ois, &c.Token)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(o.ID))

	resourceOrderRead(ctx, d, m)

	return diags
}

func resourceOrderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*hc.Client)

	var diags diag.Diagnostics

	orderID := d.Id()

	order, err := c.GetOrder(orderID, &c.Token)
	if err != nil {
		return diag.FromErr(err)
	}

	orderItems := flattenOrderItems(&order.Items)
	if err := d.Set("items", orderItems); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceOrderUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}

func resourceOrderDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}

func flattenOrderItems(orderItems *[]hc.OrderItem) []interface{} {
	if orderItems != nil {
		ois := make([]interface{}, len(*orderItems))

		for i, orderItem := range *orderItems {
			oi := make(map[string]interface{})
			oi["coffee"] = flattenCoffee(orderItem.Coffee)
			oi["quantity"] = orderItem.Quantity
			ois[i] = oi
		}
		return ois
	}

	return make([]interface{}, 0)
}

func flattenCoffee(coffee hc.Coffee) []interface{} {
	c := make(map[string]interface{})
	c["id"] = coffee.ID
	c["name"] = coffee.Name
	c["teaser"] = coffee.Teaser
	c["description"] = coffee.Description
	c["price"] = coffee.Price
	c["image"] = coffee.Image

	return []interface{}{c}
}
