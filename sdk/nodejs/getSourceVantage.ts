// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * SourceVantage DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceVantage = airbyte.getSourceVantage({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceVantage(args: GetSourceVantageArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceVantageResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceVantage:getSourceVantage", {
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceVantage.
 */
export interface GetSourceVantageArgs {
    sourceId: string;
}

/**
 * A collection of values returned by getSourceVantage.
 */
export interface GetSourceVantageResult {
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
 * SourceVantage DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceVantage = airbyte.getSourceVantage({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceVantageOutput(args: GetSourceVantageOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceVantageResult> {
    return pulumi.output(args).apply((a: any) => getSourceVantage(a, opts))
}

/**
 * A collection of arguments for invoking getSourceVantage.
 */
export interface GetSourceVantageOutputArgs {
    sourceId: pulumi.Input<string>;
}