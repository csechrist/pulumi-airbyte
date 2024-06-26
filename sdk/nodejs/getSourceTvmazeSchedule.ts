// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * SourceTvmazeSchedule DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceTvmazeschedule = airbyte.getSourceTvmazeSchedule({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceTvmazeSchedule(args: GetSourceTvmazeScheduleArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceTvmazeScheduleResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceTvmazeSchedule:getSourceTvmazeSchedule", {
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceTvmazeSchedule.
 */
export interface GetSourceTvmazeScheduleArgs {
    sourceId: string;
}

/**
 * A collection of values returned by getSourceTvmazeSchedule.
 */
export interface GetSourceTvmazeScheduleResult {
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
 * SourceTvmazeSchedule DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceTvmazeschedule = airbyte.getSourceTvmazeSchedule({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceTvmazeScheduleOutput(args: GetSourceTvmazeScheduleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceTvmazeScheduleResult> {
    return pulumi.output(args).apply((a: any) => getSourceTvmazeSchedule(a, opts))
}

/**
 * A collection of arguments for invoking getSourceTvmazeSchedule.
 */
export interface GetSourceTvmazeScheduleOutputArgs {
    sourceId: pulumi.Input<string>;
}
