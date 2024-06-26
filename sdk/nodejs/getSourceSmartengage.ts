// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * SourceSmartengage DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceSmartengage = airbyte.getSourceSmartengage({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceSmartengage(args: GetSourceSmartengageArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceSmartengageResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceSmartengage:getSourceSmartengage", {
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceSmartengage.
 */
export interface GetSourceSmartengageArgs {
    sourceId: string;
}

/**
 * A collection of values returned by getSourceSmartengage.
 */
export interface GetSourceSmartengageResult {
    /**
     * The values required to configure the source. Parsed as JSON.
     */
    readonly configuration: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly sourceId: string;
    readonly sourceType: string;
    readonly workspaceId: string;
}
/**
 * SourceSmartengage DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceSmartengage = airbyte.getSourceSmartengage({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceSmartengageOutput(args: GetSourceSmartengageOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceSmartengageResult> {
    return pulumi.output(args).apply((a: any) => getSourceSmartengage(a, opts))
}

/**
 * A collection of arguments for invoking getSourceSmartengage.
 */
export interface GetSourceSmartengageOutputArgs {
    sourceId: pulumi.Input<string>;
}
