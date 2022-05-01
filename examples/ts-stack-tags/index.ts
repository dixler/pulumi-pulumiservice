import * as pulumi from "@pulumi/pulumi";
import * as service from "@pulumi/pulumiservice";


const config = new pulumi.Config()

// export tags so you can easily use them in another stack via a stack reference
export const tags = {
    "department": "r&d",
    "cost-center": "infrastructure",
    "team": "service team"
}

// Loop through object's entries and create a StackTag resource for each entry
// pulumi.getProject() and pulumi.getStack() are defined at runtime, so no need to define it in config or hardcode
// this allows us to be very flexible about where we are deploying these tags
Object.entries(tags).forEach(([name, value]) => {
    new service.StackTag(`tag-${name}`, {
        // there is no way to look up the target organization from the pulumi SDK currently
        // so we must either define that in configuration or hardcode it
        organization: config.require("orgName"),
        project: pulumi.getProject(),
        stack: pulumi.getStack(),
        name: name,
        value: value,
    })
})
